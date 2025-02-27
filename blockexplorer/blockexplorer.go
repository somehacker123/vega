// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package blockexplorer

import (
	"context"
	"fmt"
	"net"

	"code.vegaprotocol.io/vega/blockexplorer/api"
	"code.vegaprotocol.io/vega/blockexplorer/store"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	ourGrpc "code.vegaprotocol.io/vega/blockexplorer/api/grpc"
	"code.vegaprotocol.io/vega/blockexplorer/config"
	"code.vegaprotocol.io/vega/libs/net/pipe"
	"code.vegaprotocol.io/vega/logging"
	pb "code.vegaprotocol.io/vega/protos/blockexplorer/api/v1"
)

type gatewayService interface {
	api.GatewayHandler
	Start() error
}

type grpcConnector interface {
	net.Listener
	DialGRPC(opts ...grpc.DialOption) (*grpc.ClientConn, error)
	Dial(ctx context.Context, target string) (net.Conn, error)
}

type server interface {
	Serve(net.Listener) error
}

type gateway interface {
	server
	Register(api.GatewayHandler, string)
}

type portal interface {
	Serve() error
	GRPCListener() net.Listener
	GatewayListener() net.Listener
}

type BlockExplorer struct {
	config                  config.Config
	log                     *logging.Logger
	store                   *store.Store
	blockExplorerGrpcServer pb.BlockExplorerServiceServer
	grpcServer              server
	grpcPipeConn            grpcConnector
	grpcUI                  gatewayService
	restAPI                 gatewayService
	portal                  portal
	gateway                 gateway
}

func NewFromConfig(config config.Config) *BlockExplorer {
	a := &BlockExplorer{}
	a.config = config
	a.log = logging.NewLoggerFromConfig(config.Logging)
	a.store = store.MustNewStore(config.Store, a.log)

	// main grpc api
	a.blockExplorerGrpcServer = ourGrpc.NewBlockExplorerAPI(a.store, config.API.GRPC, a.log)
	a.grpcServer = ourGrpc.NewServer(config.API.GRPC, a.log, a.blockExplorerGrpcServer)

	// grpc-ui; a web front end that talks to the grpc api through a 'pipe' (fake connection)
	a.grpcPipeConn = pipe.NewPipe("grpc-pipe")
	a.grpcUI = api.NewGRPCUIHandler(a.log, a.grpcPipeConn, config.API.GRPCUI)

	// a REST api that proxies to the GRPC api, generated by grpc-rest
	// a.restApiConn = pipe.NewPipe("rest-api")
	a.restAPI = api.NewRESTHandler(a.log, a.grpcPipeConn, config.API.REST)

	// The gateway collects all the HTTP handlers into a big 'serveMux'
	a.gateway = api.NewGateway(a.log, config.API.Gateway)
	a.gateway.Register(a.grpcUI, config.API.GRPCUI.Endpoint)
	a.gateway.Register(a.restAPI, config.API.REST.Endpoint)

	// However GRPC is special, because it uses HTTP2 and really wants to be in control
	// of its own connection. Fortunately there's a tool called cMux which creates dummy listeners
	// and peeks at the stream to decide where to send it. If it's http/2 - send to grpc server
	// otherwise dispatch to the gateway, which then sends it to which ever handler has registered
	a.portal = api.NewPortal(config.API, a.log)
	return a
}

func (a *BlockExplorer) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Two grpc services; one for internal connections using a fast pipe, and another for external
	g.Go(func() error { return a.grpcServer.Serve(a.grpcPipeConn) })
	g.Go(func() error { return a.grpcServer.Serve(a.portal.GRPCListener()) })

	// Then start our gateway and portal servers
	g.Go(func() error { return a.gateway.Serve(a.portal.GatewayListener()) })
	g.Go(func() error { return a.portal.Serve() })

	// Now we can do all the http 'handlers' that talk to the gateway
	if err := a.grpcUI.Start(); err != nil {
		return fmt.Errorf("starting grpc-ui: %w", err)
	}

	if err := a.restAPI.Start(); err != nil {
		return fmt.Errorf("starting rest grpc proxy: %w", err)
	}

	// Lastly try and gracefully shutdown if the context is cancelled
	g.Go(func() error {
		<-ctx.Done()
		return a.Close()
	})

	return g.Wait()
}

func (a *BlockExplorer) Close() error {
	// TODO
	return nil
}
