package api

import (
	"context"

	cli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/gw/v1/api"
	pb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func (h *Handler) GetAPIs(ctx context.Context) (*npool.GetAPIsResponse, error) {
	infos, total, err := cli.GetAPIs(ctx, &pb.Conds{
		Exported: &basetypes.BoolVal{
			Op:    cruder.EQ,
			Value: *h.Exported,
		},
		Depracated: &basetypes.BoolVal{
			Op:    cruder.EQ,
			Value: *h.Deprecated,
		},
		ServiceName: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: *h.ServiceName,
		},
	}, h.Offset, h.Limit)
	if err != nil {
		return nil, err
	}

	return &npool.GetAPIsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (h *Handler) GetDomains(ctx context.Context) (*npool.GetDomainsResponse, error) {
	infos, err := cli.GetDomains(ctx)
	if err != nil {
		return nil, err
	}

	return &npool.GetDomainsResponse{
		Infos: infos,
	}, nil
}
