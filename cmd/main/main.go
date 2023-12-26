package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nyan/api-crud/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusMovedPermanently)
	})
	r.HandleFunc("/hello", handler.HelloHandler)
	r.HandleFunc("/movies/{ID}", handler.GetMovie)
	r.HandleFunc("/movies", handler.GetAllMovies)
	r.HandleFunc("/addmovies", handler.AddMovie)
	// r.HandleFunc("/movies/{ID}", updateMovie)
	r.HandleFunc("/delmovies/{ID}", handler.DeleteMovie)
	http.Handle("/", r)
	fmt.Println("Starting server at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
