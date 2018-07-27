package restproxy

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net/http"
	"vega/api"
	"vega/log"
)

type restProxyServer struct{}

func NewRestProxyServer() *restProxyServer {
	return &restProxyServer{}
}

func (s *restProxyServer) Start() {
	var port = 3003
	var addr = fmt.Sprintf(":%d", port)
	log.Infof("Starting REST<>GRPC based HTTP server on port %d...\n", port)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	endpoint := "localhost:3002"
	jsonpb := &JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		// This is necessary to get error details properly marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterTradingHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalf("Registering trading handler for rest proxy endpoints %+v", err)
	} else {
		// CORS support
		handler := cors.Default().Handler(mux)
		// Gzip encoding support
		handler = NewGzipHandler(handler.(http.HandlerFunc))
		// Start http server on port specified
		http.ListenAndServe(addr, handler)
	}
}