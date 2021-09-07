package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

func (c *Client) Definition(input string) ([]protocol.Location, error) {
	path, line, char, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	definitionParams := &protocol.DefinitionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri.File(path)},
			Position: protocol.Position{
				Line:      uint32(line),
				Character: uint32(char),
			},
		},
	}

	return c.ServerDispatcher.Definition(context.TODO(), definitionParams)
}

func parseInput(input string) (string, int, int, error) {
	split := strings.Split(input, ":")

	if len(split) != 3 {
		return "", 0, 0, fmt.Errorf("failed to parse position from input: %q", input)
	}

	path := split[0]
	line, err := strconv.Atoi(split[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("error converting line position to int: %v", err)
	}

	char, err := strconv.Atoi(split[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("error converting character position to int: %v", err)
	}

	return path, line, char, nil
}
