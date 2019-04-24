package grpc

import (
	"fmt"
	"net"

	"code.vegaprotocol.io/vega/internal/parties"

	"code.vegaprotocol.io/vega/internal"
	"code.vegaprotocol.io/vega/internal/api"
	"code.vegaprotocol.io/vega/internal/blockchain"
	"code.vegaprotocol.io/vega/internal/candles"
	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/markets"
	"code.vegaprotocol.io/vega/internal/monitoring"
	"code.vegaprotocol.io/vega/internal/orders"
	"code.vegaprotocol.io/vega/internal/trades"
	"code.vegaprotocol.io/vega/internal/vegatime"

	"google.golang.org/grpc"
)

const (
	namedLogger = "api.grpc"
)

type grpcServer struct {
	log *logging.Logger
	api.Config
	stats         *internal.Stats
	client        *blockchain.Client
	orderService  *orders.Svc
	tradeService  *trades.Svc
	candleService *candles.Svc
	marketService *markets.Svc
	partyService  *parties.Svc
	timeService   *vegatime.Svc
	srv           *grpc.Server
	statusChecker *monitoring.Status
}

func NewGRPCServer(
	log *logging.Logger,
	config api.Config,
	stats *internal.Stats,
	client *blockchain.Client,
	timeService *vegatime.Svc,
	marketService *markets.Svc,
	partyService *parties.Svc,
	orderService *orders.Svc,
	tradeService *trades.Svc,
	candleService *candles.Svc,
	statusChecker *monitoring.Status,
) *grpcServer {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())

	return &grpcServer{
		log:           log,
		Config:        config,
		stats:         stats,
		client:        client,
		orderService:  orderService,
		tradeService:  tradeService,
		candleService: candleService,
		timeService:   timeService,
		marketService: marketService,
		partyService:  partyService,
		statusChecker: statusChecker,
	}
}

func (s *grpcServer) ReloadConf(cfg api.Config) {
	s.log.Info("reloading configuration")
	if s.log.GetLevel() != cfg.Level.Get() {
		s.log.Info("updating log level",
			logging.String("old", s.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		s.log.SetLevel(cfg.Level.Get())
	}

	// TODO(): not updating the the actual server for now, may need to look at this later
	// e.g restart the http server on another port or whatever
	s.Config = cfg
}

func (g *grpcServer) Start() {

	ip := g.GrpcServerIpAddress
	port := g.GrpcServerPort

	g.log.Info("Starting gRPC based API", logging.String("addr", ip), logging.Int("port", port))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		g.log.Panic("Failure listening on gRPC port", logging.Int("port", port), logging.Error(err))
	}

	var handlers = &Handlers{
		Stats:         g.stats,
		Client:        g.client,
		OrderService:  g.orderService,
		TradeService:  g.tradeService,
		CandleService: g.candleService,
		MarketService: g.marketService,
		PartyService:  g.partyService,
		TimeService:   g.timeService,
		statusChecker: g.statusChecker,
	}
	g.srv = grpc.NewServer()
	api.RegisterTradingServer(g.srv, handlers)
	err = g.srv.Serve(lis)
	if err != nil {
		g.log.Panic("Failure serving gRPC API", logging.Error(err))
	}
}

func (g *grpcServer) Stop() {
	if g.srv != nil {
		g.log.Info("Stopping gRPC based API")
		g.srv.GracefulStop()
	}
}
