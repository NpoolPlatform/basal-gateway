//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api1 "github.com/NpoolPlatform/basal-gateway/pkg/api"
	npool "github.com/NpoolPlatform/message/npool/basal/gw/v1/api"
)

func (s *Server) UpdateAPI(ctx context.Context, in *npool.UpdateAPIRequest) (*npool.UpdateAPIResponse, error) {
	handler, err := api1.NewHandler(ctx,
		api1.WithID(&in.ID, true),
		api1.WithDeprecated(in.Deprecated, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateAPI(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return info, nil
}
