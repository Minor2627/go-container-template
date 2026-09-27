package main

import (
	"log/slog"
	"net/http"

	"inholland.nl/go-container-template/internal/config"
)

const DefaultAddr string = ":8080"

func main() {
	cfg := config.NewEnvConfigProvider()
	addr := cfg.GetConfig("ADDR", DefaultAddr)

	slog.Info("Starting HTTP server", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("HTTP server error", "error", err.Error())
	}
}
