package main

import (
	"context"
	"log"
)

func main() {
	if err := Execute(context.Background()); err != nil {
		log.Fatalf("execute: %v", err)
	}
}
