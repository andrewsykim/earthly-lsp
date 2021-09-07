module earthly-lsp.dev

go 1.16

require (
	github.com/earthly/earthly v0.5.23
	go.lsp.dev/jsonrpc2 v0.9.0
	go.lsp.dev/protocol v0.11.2
	go.lsp.dev/uri v0.3.0
	go.uber.org/zap v1.19.0
)

replace (
	// estargz: needs this replace because stargz-snapshotter git repo has two go.mod modules.
	github.com/containerd/stargz-snapshotter/estargz => github.com/containerd/stargz-snapshotter/estargz v0.0.0-20201217071531-2b97b583765b
	github.com/docker/docker => github.com/docker/docker v20.10.3-0.20210609100121-ef4d47340142+incompatible

	github.com/jessevdk/go-flags => github.com/alexcb/go-flags v0.0.0-20210722203016-f11d7ecb5ee5

	github.com/moby/buildkit => github.com/earthly/buildkit v0.7.1-0.20210728003253-199ad6a5c213
	github.com/tonistiigi/fsutil => github.com/earthly/fsutil v0.0.0-20210609160335-a94814c540b2
)
