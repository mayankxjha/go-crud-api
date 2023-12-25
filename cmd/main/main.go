package main

import (
	// "encoding/json"
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	// "reflect"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)
type Movie struct{
	ID int
	Title string
	ISBN string
	Director Director
}

type Director struct{
	Fname string
	Lname string
}
func populateMovies(path string) (data []uint8){
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
func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello there, General Kenobi")
}
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	data := populateMovies("movies.txt")
	for _, movie :=range(data){
		fmt.Fprintf(w, "%s", string(movie))
	}
}
func main(){
	var movies []Movie
	err := json.Unmarshal(populateMovies("movies.txt"), &movies)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusMovedPermanently)
	})
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/movies", getAllMovies)
	// r.HandleFunc("/movies", addMovie)
	// r.HandleFunc("/movies/{ID}", getMovie)
	// r.HandleFunc("/movies/{ID}", updateMovie)
	// r.HandleFunc("/movies/{ID}", deleteMovie)
	http.Handle("/", r)
	fmt.Println("Starting server at port 8080")
	if err:=http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	}
}