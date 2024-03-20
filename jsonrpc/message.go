package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) (string, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content), nil
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (method string, content []byte, err error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find the separator")
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}
