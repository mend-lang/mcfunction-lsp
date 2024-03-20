package main

import (
	"bufio"
	"os"

	"github.com/mend-lang/mcfunction-lsp/cli"
	"github.com/mend-lang/mcfunction-lsp/jsonrpc"
	"github.com/mend-lang/mcfunction-lsp/lsp"
)

func main() {
	cli.ParseFlags()
	logger := cli.GetLogger("LSP", cli.Flags.LogFile)
	logger.Println("Program started")

	handler := lsp.Handler{Logger: logger, Contents: nil}
	var methodHandlers = map[string]func() any{
		"initialize": func() any {
			return handler.Initialize()
		},
		"initialized": func() any {
			return nil // Ignore
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(jsonrpc.Split)
	writer := os.Stdout

	for scanner.Scan() {
		method, contents, err := jsonrpc.DecodeMessage(scanner.Bytes())
		if err != nil {
			logger.Printf("Error decoding message: %s", err)
			continue
		}

		methodHandler := methodHandlers[method]
		if methodHandler == nil {
			logger.Printf("Skipping unknown method: %s", method)
			continue
		}

		handler.Contents = contents
		response := methodHandler()
		if response == nil {
			continue
		}

		msg, err := jsonrpc.EncodeMessage(response)
		if err != nil {
			logger.Printf("Failed to encode '%s': %s", method, err)
		}
		writer.Write([]byte(msg))
	}
}
