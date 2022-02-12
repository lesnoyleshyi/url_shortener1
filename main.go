package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"url_shortener1/handlers"
	"url_shortener1/storage"
)

func main() {
	var conf storage.DBconfig
	conn, err := storage.NewConnection(*conf.ParseFromEnv())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	router := handlers.NewRouter("/", conn)

	go func() {
		log.Printf("Http server start listening on port 8080\n")
		if err := http.ListenAndServe(":8080", router); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Server stopped")
			}
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)
	sig := <-sigquit
	log.Printf("%+v signal caught, shutting down server", sig)
}
