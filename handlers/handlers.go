package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ReqController struct{}

type Url struct {
	Url string `json:"url"`
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
		fmt.Fprint(w, "Provide short URL, please\n")
		return
	}
	if err := json.Unmarshal(body, &UrlFromReq); err != nil {
		fmt.Fprintf(w, "Error unmarshalling recieved url: %s\n", err)
		return
	}
	fmt.Fprintf(w, "I'll get short URL and return original\n")
	fmt.Fprintf(w, "I only can print received url yet:\n%s\n", UrlFromReq.Url)

}

func (req ReqController) GetShortRetLong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var urlFromReq Url

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if body == nil || len(body) == 0 {
		fmt.Fprint(w, "Provide URL to be shortened, please\n")
		return
	}
	if err := json.Unmarshal(body, &urlFromReq); err != nil {
		fmt.Fprintf(w, "Error unmarshalling recieved url: %s\n", err)
		return
	}

	fmt.Fprintf(w, "it is the unmarshalled url i've recieved:\n%s\n", urlFromReq.Url)
}
