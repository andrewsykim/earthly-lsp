FROM golang:1.16
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
    SAVE IMAGE --push earthly/earthly-lsp:v0.1
