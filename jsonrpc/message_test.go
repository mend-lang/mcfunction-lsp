package jsonrpc_test

import (
	"testing"

	"github.com/mend-lang/mcfunction-lsp/jsonrpc"
)

type EncodeExample struct {
	Method string `json:"method"`
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"
	actual, err := jsonrpc.EncodeMessage(EncodeExample{Method: "hi"})
	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Fatalf("Expected %s, Got %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	msg := "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"
	method, content, err := jsonrpc.DecodeMessage([]byte(msg))
	if err != nil {
		t.Fatal(err)
	}

	if len(content) != 15 {
		t.Fatalf("Expected Content-Length to be 15, got: %d", len(content))
	}

	if method != "hi" {
		t.Fatalf("Expected method to be 'hi', got: %s", method)
	}
}
