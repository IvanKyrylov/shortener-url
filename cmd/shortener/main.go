package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/IvanKyrylov/shortener-url/internal/handler"
	"github.com/IvanKyrylov/shortener-url/store"

	"github.com/IvanKyrylov/shortener-url/config"
)

func main() {
	configPath := flag.String("config", "./config/config.json", "path of the config file")
	flag.Parse()

	config, err := config.FromConfigFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := store.New(config)

	if err != nil {
		log.Fatal(err)
	}

	defer svc.Close()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler: handler.New(config.Options.Prefix, svc),
	}

	go func() {
		log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatalf("%v", err)
		} else {
			log.Println("Server closed!")
		}
	}()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)

	sig := <-sigquit
	log.Printf("caught sig: %+v", sig)
	log.Printf("Gracefully shutting down server...")
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Printf("Unable to shut down server: %v", err)
	} else {
		log.Println("Server stopped")
	}
}
