package client

import (
	"context"

	"go.lsp.dev/protocol"
)

///////////////////////////////////
////   NOT IMPLEMENTED YET
//////////////////////////////////

func (c *Client) Progress(ctx context.Context, params *protocol.ProgressParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) WorkDoneProgressCreate(ctx context.Context, params *protocol.WorkDoneProgressCreateParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) LogMessage(ctx context.Context, params *protocol.LogMessageParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) PublishDiagnostics(ctx context.Context, params *protocol.PublishDiagnosticsParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) ShowMessage(ctx context.Context, params *protocol.ShowMessageParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) ShowMessageRequest(ctx context.Context, params *protocol.ShowMessageRequestParams) (result *protocol.MessageActionItem, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) Telemetry(ctx context.Context, params interface{}) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) RegisterCapability(ctx context.Context, params *protocol.RegistrationParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) UnregisterCapability(ctx context.Context, params *protocol.UnregistrationParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) ApplyEdit(ctx context.Context, params *protocol.ApplyWorkspaceEditParams) (result bool, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) Configuration(ctx context.Context, params *protocol.ConfigurationParams) (result []interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) WorkspaceFolders(ctx context.Context) (result []protocol.WorkspaceFolder, err error) {
	panic("not implemented") // TODO: Implement
}
