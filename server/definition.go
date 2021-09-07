package server

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"

	"github.com/earthly/earthly/ast"
	"github.com/earthly/earthly/ast/spec"
)

var (
	targetRegex = `\+[a-zA-Z]+`
)

// Definition returns the location of a target definition if the provided parameters is referencing an Earthfile target.
// The current implementation has many known limitations at the moment, such as:
//   - Targets referenced from other Earthfiles is not supported yet
//   - Only Earthfile command statements are checked. With, If and For statements are ignored.
//   - There is no caching of Earthfile ASTs, so Earthfiles are parsed and reloaded for every request
func (s *Server) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location, err error) {
	uri := params.TextDocumentPositionParams.TextDocument.URI

	path := parsePathFromURI(uri)
	earthFile, err := ast.Parse(ctx, path, true)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", jsonrpc2.ErrUnknown, err)
	}

	definitionLine := int(params.TextDocumentPositionParams.Position.Line)
	definitionCol := int(params.TextDocumentPositionParams.Position.Character)
	if definitionLine < 0 || definitionCol < 0 {
		return nil, fmt.Errorf("%v: %s", jsonrpc2.ErrInvalidParams, "definition line or coloumn was negative")
	}

	targetName, err := targetForPosition(earthFile, definitionLine, definitionCol)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", jsonrpc2.ErrUnknown, err)
	}

	var locs []protocol.Location
	for _, target := range earthFile.Targets {
		if target.Name != targetName {
			continue
		}

		locs = []protocol.Location{
			{
				// TODO: support Earthfile different from requested one
				URI: uri,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      uint32(target.SourceLocation.StartLine),
						Character: uint32(target.SourceLocation.StartColumn),
					},
					End: protocol.Position{
						Line:      uint32(target.SourceLocation.EndLine),
						Character: uint32(target.SourceLocation.EndColumn),
					},
				},
			},
		}
	}

	return locs, nil
}

// targetForPosition returns a target name if the provided position is referencing a target.
func targetForPosition(earthFile spec.Earthfile, line, col int) (string, error) {
	var targetName string
	for _, target := range earthFile.Targets {
		startLine := target.SourceLocation.StartLine
		endLine := target.SourceLocation.EndLine

		if line < startLine || line > endLine {
			continue
		}

		for _, statement := range target.Recipe {
			statementLine := statement.SourceLocation.StartLine
			if line != statementLine {
				continue
			}

			var name string
			var args []string
			var sourceLocation *spec.SourceLocation
			if statement.Command != nil {
				sourceLocation = statement.Command.SourceLocation
				name = statement.Command.Name
				args = statement.Command.Args
			} else {
				continue
			}

			colStart := sourceLocation.StartColumn
			colEnd := colStart + sourceLocation.EndColumn
			if col < colStart || col > colEnd {
				continue
			}

			columnOffset := colStart
			lineContent := name + " " + strings.Join(args, " ")

			r := regexp.MustCompile(targetRegex)
			targetIndexes := r.FindAllStringIndex(lineContent, -1)
			if targetIndexes == nil {
				return "", errors.New("no matches found")
			}

			definitionIndex := col - columnOffset
			for _, targetIndex := range targetIndexes {
				startIndex := targetIndex[0]
				endIndex := targetIndex[1]

				if definitionIndex < startIndex || definitionIndex > endIndex {
					continue
				}

				targetName = parseTarget(lineContent[startIndex:endIndex])
				if targetName == "" {
					return "", errors.New("unable to parse target name")
				}

				return targetName, nil
			}
		}
	}

	return "", errors.New("no target was found")
}

// parseTarget extracts the target name if one exists in a given statement argument
func parseTarget(arg string) string {
	r := regexp.MustCompile(targetRegex)
	target := r.FindString(arg)
	if target == "" {
		return ""
	}

	return strings.TrimPrefix(target, "+")
}

// parsePathFromURI extacts the file path from a protcol.DocumentURI
func parsePathFromURI(u protocol.DocumentURI) string {
	return strings.Split(u.Filename(), ":")[0]
}
