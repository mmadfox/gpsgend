package client

import (
	"github.com/valyala/fasthttp"
)

type Client struct {
	addr    string
	httpCli *fasthttp.Client
}

func New(addr string) *Client {
	return &Client{
		addr: addr,
		httpCli: &fasthttp.Client{
			Name: "gpsgend",
		},
	}
}
