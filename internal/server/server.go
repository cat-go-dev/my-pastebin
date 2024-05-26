package server

import (
	"context"
	"fmt"
	"my-pastebin/internal/services/pasta"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	pastaService *pasta.PastaService
}

func New(pastaService *pasta.PastaService) *Server {
	return &Server{
		pastaService: pastaService,
	}
}

type StoreBody struct {
	Pasta string `json:"pasta"`
}

func (s *Server) Start(ctx context.Context) {
	r := gin.Default()

	r.GET("api/pasta/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "all",
		})
	})

	r.GET("api/pasta/:hash", func(c *gin.Context) {
		// hash := c.Param("hash")

		c.JSON(200, gin.H{
			"message": "all",
		})
	})

	r.POST("api/pasta", func(c *gin.Context) {
		body := StoreBody{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		pasta, err := s.pastaService.Store(body.Pasta)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{
				"hash":       pasta.Hash,
				"paste":      pasta.Pasta,
				"created_at": pasta.CreatedAt,
			})
		}
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println("Server starting error")
	}
}
