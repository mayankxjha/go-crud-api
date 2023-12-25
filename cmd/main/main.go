package main

import (
	// "encoding/json"
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"math/rand"
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
func structData(data []uint8) (movies []Movie){
	err := json.Unmarshal(data, &movies)
	if err != nil {
		return
	}
	return movies
}
func jsonData(data Movie) (jsonString string){
	d, err := json.Marshal(data)
	if err != nil {
		jsonString = fmt.Sprintln("Error marshaling to JSON:", err)
		return
	}
	jsonString = string(d)
	return jsonString
}
func keyGen(params map[string]string) (key int){
	key, _ = strconv.Atoi(params["ID"])
	key-=1
	return key
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello there, General Kenobi")
}
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := structData(populateMovies("movies.txt"))
	for _, movie := range(movies){
		fmt.Fprintf(w, "%s\n\n", jsonData(movie))
		
	}
}

func getMovie(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	key := keyGen(params)
	movies := structData(populateMovies("movies.txt"))
	fmt.Fprintf(w, "%s", jsonData(movies[key]))
}

func addMovie(w http.ResponseWriter, r *http.Request){
	movies := structData(populateMovies("movies.txt"))
	w.Header().Set("Content-type", "application/json")
	var Mov Movie
	_ = json.NewDecoder(r.Body).Decode(&Mov)
	Mov.ID = rand.Intn(1000000)
	movies = append(movies, Mov)
	json.NewEncoder(w).Encode(Mov)
}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusMovedPermanently)
	})
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/movies", getAllMovies)
	r.HandleFunc("/movies", addMovie)
	r.HandleFunc("/movies/{ID}", getMovie)
	// r.HandleFunc("/movies/{ID}", updateMovie)
	// r.HandleFunc("/movies/{ID}", deleteMovie)
	http.Handle("/", r)
	fmt.Println("Starting server at port 8080")
	if err:=http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	}
}