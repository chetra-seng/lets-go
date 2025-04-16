package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app := &application{
		logger: logger,
	}

	flag.Parse()

	logger.Info("Starting server ", slog.Any("addr", *addr))

	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
