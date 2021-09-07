package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"testing"

	"earthly-lsp.dev/client"
	"earthly-lsp.dev/server"
)

func Test_run(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("error running server: %v", err)
	}
	defer l.Close()

	s := server.NewServer(l)
	go func() {
		if err := s.Run(); err != nil {
			panic("error running test server")
		}
	}()

	addr := l.Addr().String()
	c := client.NewClient(addr)
	if err := c.Dial(); err != nil {
		t.Fatalf("error dialing to server: %v", err)
	}

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
    COPY . .
    RUN go test ./...

docker:
    COPY +build/earthly-lsp .
    ENTRYPOINT ["/earthly-lsp/earthly-lsp"]
    SAVE IMAGE --push earthly/earthly-lsp:v0.1`

	earthFile, err := ioutil.TempFile("", "Earthfile")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer os.Remove(earthFile.Name())

	if _, err := earthFile.Write([]byte(earthFileContent)); err != nil {
		t.Fatalf("error writing to temp file: %v", err)
	}

	results := &bytes.Buffer{}
	if err := run(c, results, []string{"definition", earthFile.Name() + ":18:11"}); err != nil {
		t.Errorf("error running earthly-lsp: %v", err)
	}

	expectedResults := fmt.Sprintf(`[{"uri":"file://%s","range":{"start":{"line":4,"character":0},"end":{"line":9,"character":0}}}]`, earthFile.Name())

	if strings.TrimSpace(results.String()) != expectedResults {
		t.Logf("actual results: %v", results.String())
		t.Logf("expected results: %v", expectedResults)
		t.Errorf("unexpected definition results")
	}
}
