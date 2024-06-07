package server

import (
	"context"
	"log/slog"
	"my-pastebin/internal/services/pasta"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	pastaService *pasta.PastaService
	logger       *slog.Logger
}

func New(pastaService *pasta.PastaService, logger *slog.Logger) *Server {
	return &Server{
		pastaService: pastaService,
		logger:       logger,
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
		collection, err := s.pastaService.GetAll(ctx)
		if err != nil {
			s.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{"result": collection})
		}
	})

	r.GET("api/pasta/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		if hash == "" {
			s.logger.Error("empty hash string")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "empty hash string",
			})
		}

		pasta, err := s.pastaService.GetByHash(ctx, hash)
		if err != nil {
			s.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{"result": pasta})
		}
	})

	r.POST("api/pasta", func(c *gin.Context) {
		body := StoreBody{}
		if err := c.BindJSON(&body); err != nil {
			s.logger.Error(err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		pasta, err := s.pastaService.Store(ctx, body.Pasta)
		if err != nil {
			s.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{"result": pasta})
		}
	})

	err := r.Run()
	if err != nil {
		s.logger.Error("Server starting error")
	}
}
