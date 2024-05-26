package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s Server) Start(ctx context.Context) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "all",
		})
	})

	r.GET("/:hash", func(c *gin.Context) {
		hash := c.Param("hash")

		c.JSON(200, gin.H{
			"message": "all",
		})
	})

	// список роутов:
	// - get all
	// - get by hash
	// - store - return hash
	// - login
	// - unlogin
	// - share

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println("Server starting error")
	}
}
