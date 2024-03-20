package lsp

import "github.com/mend-lang/mcfunction-lsp/lsp/structs"

type Response struct {
	RPC string `json:"jsonrpc"`
	Id  *int   `json:"id,omitempty"`
}

type GenericResponse[T interface{}] struct {
	Response
	Result T `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   structs.ClientInfo `json:"serverInfo"`
}
