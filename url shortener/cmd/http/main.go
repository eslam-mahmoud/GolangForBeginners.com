package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"shortener"

	"github.com/gorilla/mux"
)

func main() {
	service, err := shortener.NewService()
	if err != nil {
		fmt.Println(err)
		return
	}
	handler := &handlers{service: service}

	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.ping).Methods("GET")
	r.HandleFunc("/shorten", handler.shorten).Methods("POST")
	r.HandleFunc("/{id:[0-9a-zA-Z]+}", handler.expand).Methods("GET")

	http.Handle("/", r)
	fmt.Println("running web server on port :8090")
	http.ListenAndServe(":8090", nil)
}

type handlers struct {
	service *shortener.Service
}

func (h *handlers) ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func (h *handlers) expand(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]
	url, err := h.service.Expand(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"url": url,
	})
}

func (h *handlers) shorten(w http.ResponseWriter, req *http.Request) {
	// shortenReq define the post req params
	type shortenReq struct {
		URL string `json:"url"`
	}

	s := shortenReq{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	ID, err := h.service.Shorten(s.URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error 3": err.Error(),
		})
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"id": ID,
	})
}
