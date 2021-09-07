package server

import (
	"context"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

func Test_DefinitionSingleEarthfile(t *testing.T) {
	earthFileContent := `FROM golang:1.16
WORKDIR /earthly-lsp

deps:
    FROM +base
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    FROM +deps
    COPY . .
    RUN go build -o build/earthly-lsp
    SAVE ARTIFACT build/earthly-lsp /earthly-lsp AS LOCAL build/earthly-lsp

test:
    FROM +deps
    RUN go test ./...`

	earthFile, err := ioutil.TempFile("", "Earthfile")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer os.Remove(earthFile.Name())

	if _, err := earthFile.Write([]byte(earthFileContent)); err != nil {
		t.Fatalf("error writing to temp file: %v", err)
	}

	s := &Server{}
	params := &protocol.DefinitionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri.File(earthFile.Name())},
			Position: protocol.Position{
				Line:      uint32(12),
				Character: uint32(11),
			},
		},
	}

	results, err := s.Definition(context.TODO(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedResults := []protocol.Location{
		{
			URI: uri.File(earthFile.Name()),
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      uint32(4),
					Character: uint32(0),
				},
				End: protocol.Position{
					Line:      uint32(9),
					Character: uint32(0),
				},
			},
		},
	}

	if !reflect.DeepEqual(results, expectedResults) {
		t.Logf("actual results: %v", results)
		t.Logf("expected results: %v", expectedResults)
		t.Error("unexpected results")
	}
}

func Test_parseTarget(t *testing.T) {
	testcases := []struct {
		name   string
		arg    string
		target string
	}{
		{
			name:   "simple target arg",
			arg:    "+base",
			target: "base",
		},
		{
			name:   "target with artifact",
			arg:    "+build/earthly",
			target: "build",
		},
		{
			name:   "target in another Earthfile",
			arg:    "./ast/parser+parser/*.go",
			target: "parser",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			target := parseTarget(testcase.arg)
			if target != testcase.target {
				t.Logf("actual target: %v", target)
				t.Logf("expected target: %v", testcase.target)
				t.Errorf("unexpected target name")
			}
		})
	}
}

func Test_parsePathFromURI(t *testing.T) {
	testcases := []struct {
		name string
		arg  string
		path string
	}{
		{
			name: "simple URI",
			arg:  "file:///tmp/foo/bar",
			path: "/tmp/foo/bar",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			path := parsePathFromURI(uri.New(testcase.arg))
			if path != testcase.path {
				t.Logf("actual target path: %v", path)
				t.Logf("expected target path: %v", testcase.path)
				t.Errorf("unexpected target path")
			}
		})
	}
}
