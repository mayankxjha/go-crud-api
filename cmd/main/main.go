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
	r.HandleFunc("/addmovie", handler.AddMovie)
	r.HandleFunc("/movies/{ID}", handler.UpdateMovie)
	r.HandleFunc("/delmovie/{ID}", handler.DeleteMovie)
	http.Handle("/", r)
	fmt.Println("Starting server at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
