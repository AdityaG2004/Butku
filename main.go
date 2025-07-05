package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/AdityaG2004/butku/resolver"
	"github.com/AdityaG2004/butku/shortner"
)

func main() {
	loadenv()
	StartServer()

}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten/v1", shortner.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", resolver.ResolveURL).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(os.Getenv("SERVE_PORT"), nil))
}

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error Loading env file %v", err)
	}
}
