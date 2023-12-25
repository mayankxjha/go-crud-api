package main

import (
	// "encoding/json"
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	// "reflect"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)
type Movie struct {
	ID       int
	Title    string
	ISBN     string
	Director Director
}
type Director struct {
	Fname string
	Lname string
}
func populateMovies(path string) (data []uint8) {
	file, err := os.Open("movies.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := fileInfo.Size()
	data = make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	return
}
func populateTxt(path string, movies []Movie){
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(([]byte(jsonData(movies))))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
func structData(data []uint8) (movies []Movie) {
	err := json.Unmarshal(data, &movies)
	if err != nil {
		return
	}
	return movies
}
func jsonData(data interface{}) (jsonString string) {
	indentedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	jsonString = string(indentedData)
	return jsonString
}
func keyGen(params map[string]string) (key int) {
	key, _ = strconv.Atoi(params["ID"])
	key -= 1
	return key
}
var movies []Movie
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, General Kenobi")
}
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies = structData(populateMovies("movies.txt"))
	fmt.Fprintf(w, "%s\n\n", jsonData(movies))
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := keyGen(params)
	fmt.Fprintf(w, "%s", jsonData(movies[key]))
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var mov Movie
	_ = json.NewDecoder(r.Body).Decode(&mov)
	mov.ID = rand.Intn(1000000)
	movies = append(movies, mov)
	populateTxt("movies.txt", movies)
	json.NewEncoder(w).Encode(mov)
}
func deleteMovie(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	key := keyGen(params)
	movies = append(movies[:key], movies[key+1:]...)
	populateTxt("movies.txt", movies)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusMovedPermanently)
	})
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/movies", getAllMovies)
	r.HandleFunc("/addmovies", addMovie)
	r.HandleFunc("/movies/{ID}", getMovie)
	// r.HandleFunc("/movies/{ID}", updateMovie)
	r.HandleFunc("/movies/{ID}", deleteMovie)
	http.Handle("/", r)
	fmt.Println("Starting server at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
