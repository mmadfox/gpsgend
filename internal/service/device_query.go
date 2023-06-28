package service

import (
	"context"

	"github.com/mmadfox/gpsgend/internal/device"
)

type DeviceQuery struct {
	queryStorage device.Query
}

func NewDeviceQuery(q device.Query) *DeviceQuery {
	return &DeviceQuery{
		queryStorage: q,
	}
}

func (dq *DeviceQuery) Search(ctx context.Context, filter device.QueryFilter) (device.SearchResults, error) {
	return dq.queryStorage.Search(ctx, filter)
}
