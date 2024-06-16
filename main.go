package main

import (
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	uuidWithHyphen := uuid.New()
	response := uuidWithHyphen.String()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write([]byte(`{"uuid": "` + response + `"}`)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
