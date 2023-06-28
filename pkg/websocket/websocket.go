package websocket

import (
	"context"

	pb "github.com/mmadfox/go-gpsgen/proto"
)

func Listen(addr string, onDevice func(*pb.Device), onError func(error)) error {
	listener := NewListener()
	if onError != nil {
		listener.OnError = onError
	}
	return listener.ListenAndServe(
		context.Background(),
		addr+"/gpsgend/ws",
		HandlerFunc(onDevice))
}
