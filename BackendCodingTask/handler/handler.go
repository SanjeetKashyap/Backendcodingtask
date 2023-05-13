package handler

import (
	"Backend/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Database server details
const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "pgsql0009"
	DB_NAME     = "postgres"
)

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		fmt.Println("the Error is", err)
	}
}

// Connection to Database
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

// Handlers

func GetLongestDurationMovies(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	//Query
	rows, err := db.Query("SELECT tconst, primarytitle, runtimeminutes, genres FROM movies ORDER BY runtimeminutes DESC LIMIT 10")
	checkErr(err)

	var movies []model.Movies

	for rows.Next() {
		var tconst string

		var primaryTitle string
		var runtimeminutes int
		var genre string
		err := rows.Scan(&tconst, &primaryTitle, &runtimeminutes, &genre)

		checkErr(err)

		movies = append(movies, model.Movies{Tconst: tconst, PrimaryTitle: primaryTitle, RuntimeMinutes: runtimeminutes, Genres: genre})
	}
	// Json formating
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func CreateNewMovie(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	tConst := r.FormValue("tconst")
	titleType := r.FormValue("titletype")
	primaryTitle := r.FormValue("primarytitle")
	runtimeMinutes := r.FormValue("runtimeminutes")
	genres := r.FormValue("genres")

	averageRating := r.FormValue("averagerating")
	numVotes := r.FormValue("numvotes")

	_, err1 := db.Exec("INSERT INTO movies(tconst,titletype,primarytitle,runtimeminutes,genres) VALUES($1,$2,$3,$4,$5)", tConst, titleType, primaryTitle, runtimeMinutes, genres)
	_, err2 := db.Exec("INSERT INTO ratings(tconst,averagerating,numvotes) VALUES($1,$2,$3)", tConst, averageRating, numVotes)

	checkErr(err1)
	checkErr(err2)

	if err1 == nil && err2 == nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Success")
	}
}

func GetTopRatedMovies(w http.ResponseWriter, r *http.Request) {

	db := setupDB()
	rows, err := db.Query("SELECT movies.tconst, primarytitle, genres, averagerating FROM movies JOIN ratings ON movies.tconst = ratings.tconst WHERE averagerating > 6 ORDER BY averagerating desc")
	checkErr(err)
	var movies []model.Movies

	for rows.Next() {
		var tconst string
		var primaryTitle string
		var genre string
		var averagerating float64

		err := rows.Scan(&tconst, &primaryTitle, &genre, &averagerating)
		checkErr(err)

		movies = append(movies, model.Movies{Tconst: tconst, PrimaryTitle: primaryTitle, Genres: genre, Averagerating: averagerating})

	}
	// Json formating
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetGenreMoviesWithSubtotals(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	rows, err := db.Query("SELECT genres, COALESCE(primarytitle,'TOTAL') as primarytitle  , sum(numvotes) as numvotes FROM movies JOIN ratings ON movies.tconst = ratings.tconst GROUP BY ROLLUP (genres, primarytitle) ORDER BY genres, primarytitle")
	checkErr(err)
	var movies []model.Movies

	for rows.Next() {
		var genre string
		var primaryTitle string
		var numVotes int

		err := rows.Scan(&genre, &primaryTitle, &numVotes)
		checkErr(err)
		movies = append(movies, model.Movies{Genres: genre, PrimaryTitle: primaryTitle, Numvotes: numVotes})
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func UpdateRuntimeMinutes(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	_, err := db.Exec("UPDATE movies SET runtimeminutes = CASE WHEN genres = 'Documentary' THEN runtimeminutes + 15 WHEN genres = 'Animation' THEN runtimeminutes + 30 ELSE runtimeminutes + 45 END;")

	checkErr(err)
	if err == nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Success")
	}
}
