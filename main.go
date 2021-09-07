package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"earthly-lsp.dev/client"
	"earthly-lsp.dev/server"
)

const (
	defaultAddr = "127.0.0.1:8080"
)

func main() {
	var remote string
	flag.StringVar(&remote, "remote", "", "the remote address for the Earthly language server")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected at least 1 subcommand")
		os.Exit(1)
	}

	// only start the LSP server locally if a remote address is not provided.
	if remote == "" {
		remote = defaultAddr

		l, err := net.Listen("tcp", remote)
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
	}

	c := client.NewClient(remote)
	if err := c.Dial(); err != nil {
		fmt.Printf("error connecting to server: %v\n", err)
		os.Exit(1)
	}

	if err := run(c, os.Stdout, os.Args[1:]); err != nil {
		fmt.Printf("error running earthly-lsp: %v\n", err)
		os.Exit(1)
	}
}

func run(c *client.Client, w io.Writer, args []string) error {
	if err := c.Initialize(context.TODO()); err != nil {
		return fmt.Errorf("error with initial protocol handshake with server: %w", err)
	}

	// TODO: handle more client subcommands as more of the LSP protocol spec is implemented
	switch subcommand := args[0]; subcommand {
	case "definition":
		if len(args) < 2 {
			return errors.New("expected at least 1 subcommand for definitions")
		}

		input := args[1]
		locs, err := c.Definition(input)
		if err != nil {
			return fmt.Errorf("error getting definition location: %w", err)
		}

		enc := json.NewEncoder(w)
		if err := enc.Encode(locs); err != nil {
			return fmt.Errorf("error encoding json to stdout: %w", err)
		}
	default:
		return fmt.Errorf("unexpected subcommand: %q", subcommand)
	}

	return nil
}
