package server

import (
	"context"

	"go.lsp.dev/protocol"
)

///////////////////////////////////
////   NOT IMPLEMENTED YET
//////////////////////////////////

func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Exit(ctx context.Context) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) CodeLensRefresh(ctx context.Context) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) SemanticTokensRefresh(ctx context.Context) (err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
	panic("not implemented") // TODO: Implement
}
