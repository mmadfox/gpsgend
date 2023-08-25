package broker

type Client interface {
	Send(data []byte)
	Close()
}
