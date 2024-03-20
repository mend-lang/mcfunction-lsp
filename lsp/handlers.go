package lsp

import (
	"encoding/json"
	"log"

	"github.com/mend-lang/mcfunction-lsp/cli"
	"github.com/mend-lang/mcfunction-lsp/lsp/structs"
)

type Handler struct {
	Logger   *log.Logger
	Contents []byte
}

func (handler Handler) readRequestContents(request any) {
	if err := json.Unmarshal(handler.Contents, &request); err != nil {
		handler.Logger.Printf("Failed to parse: %s", err)
	}
}

func (handler Handler) makeResponse(id int) Response {
	return Response{
		RPC: "2.0",
		Id:  &id,
	}
}

func (handler Handler) Initialize() any {
	var request GenericRequest[InitializeParams]
	handler.readRequestContents(&request)

	if cli.Flags.Verbose {
		handler.Logger.Printf(
			"Connected to: %s v%s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)
	}

	return GenericResponse[InitializeResult]{
		Response: handler.makeResponse(request.Id),
		Result: InitializeResult{
			Capabilities: NewServerCapabilities(),
			ServerInfo: structs.ClientInfo{
				Name:    "mcfunction-lsp",
				Version: "0.0.1-alpha",
			},
		},
	}
}
