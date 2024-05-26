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

	// todo: make swagger

	r.GET("api/pasta/all", func(c *gin.Context) {
		// todo: think about pagination
		collection, err := s.pastaService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{"result": collection})
		}
	})

	r.GET("api/pasta/:hash", func(c *gin.Context) {
		// todo: vavidation
		hash := c.Param("hash")
		pasta, err := s.pastaService.GetByHash(hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{"result": pasta})
		}
	})

	r.POST("api/pasta", func(c *gin.Context) {
		// todo: vavidation
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
			// todo: make by json response model
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
