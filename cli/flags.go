package cli

import (
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

var Flags struct {
	Verbose           bool   `short:"v" long:"verbose"  description:"Show verbose debug information"`
	LogFile           string `short:"l" long:"log-file" description:"Path to which log will be written"`
	DefaultPackFormat int8   `short:"f" long:"default-pack-format" description:"The default pack-format to use if no pack.mcmeta is found"`
}

func ParseFlags() {
	_, err := flags.Parse(&Flags)
	if err != nil {
		os.Exit(1)
	}

	// Log file
	if len(Flags.LogFile) == 0 {
		Flags.LogFile = filepath.Join(os.TempDir(), "mcfunction_lsp.log")
	}

	dir := filepath.Dir(Flags.LogFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}

	// Default pack format
	if Flags.DefaultPackFormat == 0 {
		Flags.DefaultPackFormat = 26 // 1.20.4
	}
}
