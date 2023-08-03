package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
