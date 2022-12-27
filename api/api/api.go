package api

import (
	"context"

	api1 "github.com/NpoolPlatform/message/npool/basal/gw/v1/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	api1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	api1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return api1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
