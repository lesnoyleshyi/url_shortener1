package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"url_shortener1/encoder"
	"url_shortener1/storage"
)

type handler struct {
	endpoint string
	storage  storage.Service
}

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"response"`
}

func NewRouter(endpoint string, storage storage.Service) *mux.Router {
	router := mux.NewRouter()
	handler := handler{endpoint: endpoint, storage: storage}
	router.HandleFunc(endpoint, respHandler(handler.ProcessShort)).Methods("GET")
	router.HandleFunc(endpoint, respHandler(handler.ProcessLong)).Methods("POST")
	return router
}

func respHandler(h func(io.Writer, *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(response{Success: err == nil, Data: data})
		if err != nil {
			log.Printf("Error encoding response: %s", err)
		}
	}
}

func (h handler) ProcessShort(w io.Writer, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodGet {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("method %s isn't allowed", r.Method)
	}
	var UrlFromReq encoder.Url

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("unable to parse url-request: %s", err)
	}
	if body == nil || len(body) == 0 {
		return nil, http.StatusBadRequest, fmt.Errorf("provide shortened url, please")
	}
	if err := json.Unmarshal(body, &UrlFromReq); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error unmarshalling recieved url: %s", err)
	}
	ret, err := h.storage.Retrieve(UrlFromReq.UrlShort)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, http.StatusBadRequest, fmt.Errorf("no such short url within storage")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("unable to retrieve url from storage: %s", err)
	}
	return ret, http.StatusCreated, nil
}

func (h handler) ProcessLong(w io.Writer, r *http.Request) (interface{}, int, error) {
	var UrlFromReq encoder.Url

	defer r.Body.Close()
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("method %s isn't allowed", r.Method)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("unable to parse url-request: %s", err)
	}
	if body == nil || len(body) == 0 {
		return nil, http.StatusBadRequest, fmt.Errorf("provide url to be shortened, please")
	}
	if err := json.Unmarshal(body, &UrlFromReq); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error unmarshalling recieved url: %s", err)
	}
	UrlFromReq.Encode()
	ret, err := h.storage.Save(UrlFromReq.UrlShort, UrlFromReq.UrlLong)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("unable to save url in storage: %s", err)
	}
	return ret, http.StatusCreated, nil
}
