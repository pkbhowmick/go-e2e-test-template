package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/hello", helloHandler)

	srv := &http.Server{
		Addr:    "8080",
		Handler: r,
	}
	srv.ListenAndServe()
}
