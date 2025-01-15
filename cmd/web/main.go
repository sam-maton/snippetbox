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
	flag.Parse()

	app := application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})),
	}

	app.logger.Info("starting server on http://localhost" + *addr)

	err := http.ListenAndServe(*addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}
