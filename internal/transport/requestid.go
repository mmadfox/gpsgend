package transport

import (
	"context"

	transportgrpc "github.com/mmadfox/gpsgend/internal/transport/grpc"
)

func RequestIDFromContext(ctx context.Context) string {
	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		return reqID
	}
	return transportgrpc.RequestIDFromContext(ctx)
}
