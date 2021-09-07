package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	"earthly-lsp.dev/client"
	"earthly-lsp.dev/server"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "", "the remote address for the Earthly language server")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected at least 1 subcommand")
		os.Exit(1)
	}

	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("error running server: %v\n", err)
		os.Exit(1)
	}
	defer l.Close()

	s := server.NewServer(l)
	go func() {
		if err := s.Run(); err != nil {
			fmt.Printf("error running server: %v\n", err)
			os.Exit(1)
		}
	}()

	c := client.NewClient(addr)
	if err := c.Dial(); err != nil {
		fmt.Printf("error connecting to server: %v\n", err)
		os.Exit(1)
	}

	// TODO: initialize server/client handshake
	if err := c.Initialize(context.TODO()); err != nil {
		fmt.Printf("error with initial protocol handshake with server: %v\n", err)
		os.Exit(1)
	}

	// TODO: handle more client subcommands as more of the LSP protocol spec is implemented
	switch subcommand := os.Args[1]; subcommand {
	case "definition":
		if len(os.Args) < 3 {
			fmt.Println("expected at least 1 subcommand for definitions")
			os.Exit(1)
		}

		input := os.Args[2]
		locs, err := c.Definition(input)
		if err != nil {
			fmt.Printf("error getting definition location: %v\n", err)
			os.Exit(1)
		}

		enc := json.NewEncoder(os.Stdout)
		if err := enc.Encode(locs); err != nil {
			fmt.Printf("error encoding json to stdout: %v", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("unexpected subcommand: %q\n", subcommand)
	}
}
