package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"nyan/api-crud/internal/models"
	"nyan/api-crud/internal/utilities"
)

var movies []mod.Movie

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, General Kenobi")
}
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies = util.StructData(util.PopulateMovies("movies.txt"))
	fmt.Fprintf(w, "%s\n\n", util.JsonData(movies))
}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := util.KeyGen(params)
	del := false
	for it, movie := range movies {
		if movie.ID == key {
			fmt.Fprintf(w, "%s", util.JsonData(movies[it]))
		}
		del = true
	}
	if !del {
		fmt.Fprintf(w, "Movie not found")
	}
}
func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var mov mod.Movie
	_ = json.NewDecoder(r.Body).Decode(&mov)
	mov.ID = rand.Intn(1000000)
	movies = append(movies, mov)
	util.PopulateTxt("movies.txt", movies)
	json.NewEncoder(w).Encode(mov)
}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := util.KeyGen(params)
	del := false
	for it, movie := range movies {
		if movie.ID == key {
			fmt.Fprintf(w, "Movie Deleted")
			movies = append(movies[:it], movies[it+1:]...)
			del = true
		}
	}
	if !del {
		fmt.Fprintf(w, "Movie not dound")
	}
	util.PopulateTxt("movies.txt", movies)
}
