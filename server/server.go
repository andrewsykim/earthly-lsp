package server

import (
	"context"
	"net"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
	"go.uber.org/zap"
)

func NewServer(listener net.Listener) *Server {
	return &Server{
		listener: listener,
	}
}

var _ protocol.Server = (*Server)(nil)

type Server struct {
	listener net.Listener
}

func (s *Server) Run() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// TODO log this error when it's not so noisy
			continue
		}

		stream := jsonrpc2.NewStream(conn)
		_, _, _ = protocol.NewServer(context.TODO(), s, stream, zap.NewNop())
	}
}

func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			DefinitionProvider: struct{}{},
		},
	}, nil
}
