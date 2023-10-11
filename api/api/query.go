//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/basal/gw/v1/api"

	api1 "github.com/NpoolPlatform/basal-gateway/pkg/api"
)

func (s *Server) GetAPIs(ctx context.Context, in *npool.GetAPIsRequest) (*npool.GetAPIsResponse, error) {
	handler, err := api1.NewHandler(ctx,
		api1.WithOffset(in.GetOffset()),
		api1.WithLimit(in.GetLimit()),
		api1.WithExported(in.Exported, false),
		api1.WithDeprecated(in.Depracated, false),
		api1.WithServiceName(in.ServiceName, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAPIs",
			"In", in,
			"Error", err,
		)
		return &npool.GetAPIsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.GetAPIs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAPIs",
			"In", in,
			"Error", err,
		)
		return &npool.GetAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return infos, nil
}

func (s *Server) GetDomains(ctx context.Context, in *npool.GetDomainsRequest) (*npool.GetDomainsResponse, error) {
	handler, err := api1.NewHandler(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDomains",
			"In", in,
			"Error", err,
		)
		return &npool.GetDomainsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	domains, err := handler.GetDomains(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDomains",
			"In", in,
			"Error", err,
		)
		return &npool.GetDomainsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return domains, nil
}
