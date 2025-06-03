package main

import (
	"context"

	"os"

	"github.com/jalsarraf0/ai-chat-cli/cmd/ai-chat/internal/cmd"
)

func main() {
	root := cmd.NewRootCmd()
	if err := root.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)

	"log"
)

func main() {
	if err := Execute(context.Background()); err != nil {
		log.Fatalf("execute: %v", err)

	}
}
