package service

import (
	"context"

	"github.com/mmadfox/gpsgend/internal/device"
	"golang.org/x/exp/slog"
)

type DeviceQuery struct {
	queryStorage device.Query
	logger       *slog.Logger
}

func NewDeviceQuery(q device.Query, logger *slog.Logger) *DeviceQuery {
	return &DeviceQuery{
		queryStorage: q,
		logger:       logger,
	}
}

func (dq *DeviceQuery) Search(ctx context.Context, filter device.QueryFilter) (device.SearchResult, error) {
	result, err := dq.queryStorage.Search(ctx, filter)
	if err != nil {
		dq.logger.Error("device search", "err", err)
		return result, err
	}
	return result, nil
}
