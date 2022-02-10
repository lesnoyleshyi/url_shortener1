package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"url_shortener1/handlers"
)

func main() {
	router := mux.NewRouter()
	req := handlers.ReqController{}

	router.HandleFunc("/", req.GetLongRetShort).Methods("GET")
	router.HandleFunc("/", req.GetShortRetLong).Methods("POST")

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("failed starting server")
	}
}
