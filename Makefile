dev:
	go build -o dist/lsp main.go

prod:
	GOOS="linux" GOARCH="amd64" go build -o dist/mcfunction-lsp-linux-amd64 -ldflags "-s -w" main.go
	GOOS="windows" GOARCH="amd64" go build -o dist/mcfunction-lsp-windows-amd64.exe -ldflags "-s -w" main.go
	GOOS="darwin" GOARCH="amd64" go build -o dist/mcfunction-lsp-darwin-amd64 -ldflags "-s -w" main.go
