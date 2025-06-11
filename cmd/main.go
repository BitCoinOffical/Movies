package main

import (
	"net/http"

	"main.go/internal/handlers"
	"main.go/internal/routers"
	"main.go/internal/storage"
)

func main() {
	db := storage.SQLiteINIT("movies.db")
	h := handlers.NewHandler(db)
	mux := http.NewServeMux()
	mux.Handle("/movies", routers.MethodRouter{
		"GET":  h.GetMoviesHandler,
		"POST": h.AddMoviesHandler,
	})
	mux.Handle("/movies/", routers.MethodRouter{
		"PUT":    h.UpdateMoviesHandler,
		"GET":    h.GetMoviesByIDHandler,
		"DELETE": h.DeleteMoviesHandler,
	})

	mux.Handle("/", http.FileServer(http.Dir("app")))
	http.ListenAndServe(":8080", mux)
}
