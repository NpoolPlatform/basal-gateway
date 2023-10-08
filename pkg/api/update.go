package api

import (
	"context"

	cli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	npool "github.com/NpoolPlatform/message/npool/basal/gw/v1/api"
	pb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func (h *Handler) UpdateAPI(ctx context.Context) (*npool.UpdateAPIResponse, error) {
	info, err := cli.UpdateAPI(ctx, &pb.APIReq{
		ID:         h.ID,
		Deprecated: h.Deprecated,
	})
	if err != nil {
		return nil, err
	}

	return &npool.UpdateAPIResponse{
		Info: info,
	}, nil
}
