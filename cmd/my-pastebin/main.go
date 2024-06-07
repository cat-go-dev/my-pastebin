package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"my-pastebin/internal/server"
	"my-pastebin/internal/services/config"
	"my-pastebin/internal/services/pasta"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Server starting")

	ctx := context.Background()

	config, err := config.Load("../../")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	dBRepository := pasta.NewDBRepository(db, logger)
	pastaService := pasta.NewPastaService(dBRepository, logger)

	fmt.Println("starting server")
	server := server.New(pastaService, logger)
	server.Start(ctx)

	// todo: graceful shutdown
}
