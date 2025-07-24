package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mohsinking2002/students-api-go-crud/internal/config"
	"github.com/mohsinking2002/students-api-go-crud/internal/http/student"
)

func main()  {
	// load config
	cfg := config.MustLoad()

	// database setup
	
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students/new", student.New())

	// setup server
	server := http.Server{
		Addr: cfg.Address,
		Handler: router,
	}
	slog.Info("Server running..", slog.String("address", server.Addr))
	//gracefull shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func ()  {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start the server!")
		}
	}()

	<- done

	//now shutdown
	slog.Info("Shutting down the server.")
	// ? what's context. (js abort controller)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully.")

}