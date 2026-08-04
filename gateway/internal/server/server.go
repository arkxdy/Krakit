package server

import (
	"github.com/gin-gonic/gin"
	"github.com/krakit/gateway/internal/config"
)

type Server struct {
	engine *gin.Engine
}

func New(cfg config.Config) *Server {
	engine := NewRouter(cfg)
	return &Server{
		engine: engine,
	}
}

func (s *Server) Run() {
	s.engine.Run(":8080")
}
