package api

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

type Handler struct {
	ID          *uint32
	Protocol    *npool.Protocol
	ServiceName *string
	Method      *npool.Method
	MethodName  *string
	Path        *string
	PathPrefix  *string
	Exported    *bool
	Deprecated  *bool
	Domains     *[]string
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithServiceName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid servicename")
			}
			return nil
		}
		const leastNameLen = 2
		if len(*name) < leastNameLen {
			return fmt.Errorf("service name %v too short", *name)
		}

		h.ServiceName = name
		return nil
	}
}

func WithExported(exported *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if exported == nil {
			if must {
				return fmt.Errorf("invalid exported")
			}
			return nil
		}
		h.Exported = exported
		return nil
	}
}

func WithDeprecated(deprecated *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deprecated == nil {
			if must {
				return fmt.Errorf("invalid deprecated")
			}
			return nil
		}
		h.Deprecated = deprecated
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
