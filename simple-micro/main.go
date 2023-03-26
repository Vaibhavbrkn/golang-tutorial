package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json/id`
	Isbn     string    `json/isbn`
	Title    string    `json/title`
	Director *Director `json/dirctor`
}

type Director struct {
	FirstName string `json/firstname`
	LastName  string `json/lastname`
}

var movies []Movie

func main() {
	fmt.Println("Hello from Simple Web Server")
	movies = append(movies, Movie{ID: "1", Isbn: "76447", Title: "RRR", Director: &Director{FirstName: "SS", LastName: "Rajamouli"}})
	movies = append(movies, Movie{ID: "2", Isbn: "85647", Title: "KGF", Director: &Director{FirstName: "Vaibhav", LastName: "Agrawal"}})
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatin/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatin/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatin/json")

	params := mux.Vars(r)

	for Index, Item := range movies {
		if Item.ID == params["id"] {
			movies = append(movies[:Index], movies[Index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatin/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatin/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(movies)
}
