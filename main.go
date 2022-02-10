package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"url_shortener1/handlers"
	"url_shortener1/storage"
)

func main() {
	router := mux.NewRouter()
	req := handlers.ReqController{}

	router.HandleFunc("/", req.GetLongRetShort).Methods("GET")
	router.HandleFunc("/", req.GetShortRetLong).Methods("POST")

	conf := storage.DBconfig{
		User:   "go_user",
		Passwd: "8246go",
		Host:   "localhost",
		Port:   "5432",
		DbName: "url_storage",
	}

	conn, err := storage.NewConnection(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
}
