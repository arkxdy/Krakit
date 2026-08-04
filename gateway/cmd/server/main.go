package main

import (
	"github.com/krakit/gateway/internal/config"
	"github.com/krakit/gateway/internal/server"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	srv.Run()

}
