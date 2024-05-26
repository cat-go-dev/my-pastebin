package main

import (
	"context"
	"fmt"

	"my-pastebin/internal/server"
)

func main() {
	// todo: parse config
	// todo: replace with logs

	ctx := context.Background()

	fmt.Println("starting server")
	server := server.New()
	server.Start(ctx)

	// todo: graceful shutdown
}
