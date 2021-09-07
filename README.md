# earthly-lsp

earthly-lsp is a starting implementation of an LSP server for Earthly.

## Examples

The earthly-lsp command currently embeds both the LSP client and server. You can use CLI subcommands to get the location of target definitions. For example:

```bash
$ earthly-lsp definition Earthfile:16:11
[{"uri":"file:///home/asykim/code/earthly-lsp/Earthfile","range":{"start":{"line":5,"character":0},"end":{"line":13,"character":0}}}]
```

## Development and Testing

Use Earthly to develop and test earthly-lsp!

Run tests:
```bash
$ earthly +test
```

Build the earthly-lsp binary:
```bash
$ earthly +build
```

Build the earthly-lsp container:
```bash
$ earthly +docker
```
