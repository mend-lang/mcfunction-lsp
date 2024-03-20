package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
type ServerCapabilities struct {
	TextDocumentSync int  `json:"textDocumentSync"`
	HoverProvider    bool `json:"hoverProvider"`
}

func NewServerCapabilities() ServerCapabilities {
	return ServerCapabilities{
		TextDocumentSync: 2, // Incremental
		HoverProvider:    true,
	}
}
