package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"my-pastebin/internal/server"
	"my-pastebin/internal/services/config"
	"my-pastebin/internal/services/pasta"
)

func main() {
	// todo: logs
	// todo: ratelimitter middleware

	ctx := context.Background()

	config, err := config.Load("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	dBRepository := pasta.NewDBRepository(db)
	pastaService := pasta.NewPastaService(dBRepository)

	fmt.Println("starting server")
	server := server.New(pastaService)
	server.Start(ctx)

	// todo: graceful shutdown
}
