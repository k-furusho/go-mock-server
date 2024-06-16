package main

import (
	"io/ioutil"
	"net/http"
)

func main()  {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request)  {
	byteArray,_ := ioutil.ReadFile("response.json")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(byteArray)
}