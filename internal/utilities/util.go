package util

import (
	"encoding/json"
	"fmt"
	"nyan/api-crud/internal/models"
	"os"
)

import (
	"strconv"
)

func KeyGen(params map[string]string) (key int) {
	key, _ = strconv.Atoi(params["ID"])
	return key
}

func StructData(data []uint8) (movies []mod.Movie) {
	err := json.Unmarshal(data, &movies)
	if err != nil {
		return
	}
	return movies
}
func JsonData(data interface{}) (jsonString string) {
	indentedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	jsonString = string(indentedData)
	return jsonString
}
func PopulateMovies(path string) (data []uint8) {
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
func PopulateTxt(path string, movies []mod.Movie) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(([]byte(JsonData(movies))))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
