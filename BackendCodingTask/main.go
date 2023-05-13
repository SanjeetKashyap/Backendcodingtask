package main

import (
	"Backend/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	// Router to perform http operations
	r := mux.NewRouter()

	//all routes 
	r.HandleFunc("/api/v1/longest-duration-movies", handler.GetLongestDurationMovies).Methods("GET")
	r.HandleFunc("/api/v1/new-movie", handler.CreateNewMovie).Methods("POST")
	r.HandleFunc("/api/v1/top-rated-movies", handler.GetTopRatedMovies).Methods("GET")
	r.HandleFunc("/api/v1/genre-movies-with-subtotals", handler.GetGenreMoviesWithSubtotals).Methods("GET")
	r.HandleFunc("/api/v1/update-runtime-minutes", handler.UpdateRuntimeMinutes).Methods("PUT")

	//here host url
	log.Fatal(http.ListenAndServe("localhost:5000", r))

}
