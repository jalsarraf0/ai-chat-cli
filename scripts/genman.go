//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/jalsarraf0/ai-chat-cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	dir := "docs/man"
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		log.Fatal(err)
	}
	if err := doc.GenManTree(cmd.NewRootCmd(), &doc.GenManHeader{}, dir); err != nil {
		log.Fatal(err)
	}
}
