package main

import (
	"log/slog"
	"net/http"

	"inholland.nl/go-container-template/internal/config"
)

const DEFAULT_ADDR string = ":8080"

func main() {
	config := config.NewEnvConfigProvider()
	addr := config.GetConfig("ADDR", DEFAULT_ADDR)

	slog.Info("Starting HTTP server", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("HTTP server error", "error", err.Error())
	}
}
