package lsp

import "github.com/mend-lang/mcfunction-lsp/lsp/structs"

type Request struct {
	RPC    string `json:"jsonrpc"`
	Id     int    `json:"id"`
	Method string `json:"method"`
}

type GenericRequest[T interface{}] struct {
	Request
	Params T `json:"params"`
}

type InitializeParams struct {
	ClientInfo *structs.ClientInfo `json:"clientInfo"`
}
