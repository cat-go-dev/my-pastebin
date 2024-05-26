package main

import (
	"context"
	"database/sql"
	"fmt"

	"my-pastebin/internal/server"
	"my-pastebin/internal/services/pasta"
)

func main() {
	// todo: config
	// todo: logs
	// todo: ratelimitter middleware

	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./database/pastebin.db")
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
