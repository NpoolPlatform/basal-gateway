package api

import (
	"context"

	basal "github.com/NpoolPlatform/message/npool/basal/gw/v1"

	api1 "github.com/NpoolPlatform/basal-gateway/api/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	basal.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	basal.RegisterGatewayServer(server, &Server{})
	api1.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := basal.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := api1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
