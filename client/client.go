package client

import (
	"context"
	"net"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

type Client struct {
	addr string

	ServerDispatcher   protocol.Server
	ServerCapabilities protocol.ServerCapabilities
}

func (c *Client) Dial() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}

	stream := jsonrpc2.NewStream(conn)
	_, _, serverDispatcher := protocol.NewClient(context.TODO(), c, stream, zap.NewNop())
	c.ServerDispatcher = serverDispatcher
	return nil
}

func (c *Client) Initialize(ctx context.Context) error {
	initParams := &protocol.InitializeParams{}
	initResults, err := c.ServerDispatcher.Initialize(context.TODO(), initParams)
	if err != nil {
		return err
	}

	c.ServerCapabilities = initResults.Capabilities
	return nil
}
