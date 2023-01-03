package main

import (
	"awesome-go-project/internal/configuration"
	"awesome-go-project/internal/handlers"
	"awesome-go-project/internal/router"
	"fmt"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"
	"time"
)

var app configuration.Application

func main() {

	initializeSlog()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Router(),
	}

	app.Logger.Info("Server started on port 8080")

	handlers.NewHandlers(&app)

	err := srv.ListenAndServe()
	if err != nil {
		app.Logger.Error("Fatal error", err)
		panic(err)
	}
}

// This is a simple logger that I use.
// More info here: https://www.youtube.com/watch?v=gd_Vyb5vEw0
func initializeSlog() {
	f, _ := os.OpenFile(fmt.Sprint("./logs/", time.Now().Format("2006-01-02"), ".log"), os.O_APPEND|os.O_CREATE|os.O_APPEND, 0666)
	wrt := io.MultiWriter(os.Stdout, f)

	th := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}.NewTextHandler(wrt)

	app.Logger = slog.New(th)
}
