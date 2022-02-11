package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"url_shortener1/storage"
)

type ReqController struct{}

type Url struct {
	UrlShort string `json:"url_short"`
	UrlLong  string `json:"url_long"`
}
type handler struct {
	endpoint string
	storage  storage.Service
}

func NewRouter(endpoint string, storage storage.Service) *mux.Router {
	router := mux.NewRouter()
	handler := handler{endpoint: endpoint, storage: storage}
	router.HandleFunc(endpoint, respHandler(handler.ProcessShort)).Methods("GET")
	router.HandleFunc(endpoint, respHandler(handler.ProcessLong)).Methods("POST")
	return router
}

func respHandler(h func() ) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request)

}

func (h handler) ProcessShort() {
	//h.storage.Save()
	return
}

func (h handler) ProcessLong() {

}

func (req ReqController) GetLongRetShort(w http.ResponseWriter, r *http.Request) {
	var UrlFromReq Url

	w.WriteHeader(http.StatusOK)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if body == nil || len(body) == 0 {
		fmt.Fprint(w, "Provide URL to be shortened, please\n")
		return
	}
	if err := json.Unmarshal(body, &UrlFromReq); err != nil {
		fmt.Fprintf(w, "Error unmarshalling recieved url: %s\n", err)
		return
	}
	fmt.Fprintf(w, "I'll get long URL and return shortened\n")
	fmt.Fprintf(w, "I only can print received url yet:\n%s\n", UrlFromReq.UrlLong)

}

func (req ReqController) GetShortRetLong(w http.ResponseWriter, r *http.Request) {
	var UrlFromReq Url
	w.WriteHeader(http.StatusOK)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if body == nil || len(body) == 0 {
		fmt.Fprint(w, "Provide shortened URL, please\n")
		return
	}
	if err := json.Unmarshal(body, &UrlFromReq); err != nil {
		fmt.Fprintf(w, "Error unmarshalling recieved url: %s\n", err)
		return
	}

	fmt.Fprintf(w, "I'll get shortened URL and return original\n")
	fmt.Fprintf(w, "I only can print received url yet:\n%s\n", UrlFromReq.UrlShort)
}
