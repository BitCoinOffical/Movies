package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"main.go/internal/models"
	"main.go/internal/storage"
)

type Handler struct {
	data *storage.DataBase
}

func NewHandler(db *storage.DataBase) *Handler {
	return &Handler{data: db}
}
func (h *Handler) AddMoviesHandler(w http.ResponseWriter, r *http.Request) {
	var movies models.Movies
	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	err = h.data.AddMovies(&movies)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movies)
}

func (h *Handler) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := h.data.GetMovies()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	if movies == nil {
		movies = []models.Movies{}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func (h *Handler) GetMoviesByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	IdMovies := strings.TrimPrefix(r.URL.Path, "/movies/")
	id, _ := strconv.Atoi(IdMovies)
	movies, err := h.data.GetMoviesByID(id)
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	if movies == nil {
		movies = &models.Movies{}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	IdMovies := strings.TrimPrefix(r.URL.Path, "/movies/")
	id, _ := strconv.Atoi(IdMovies)
	h.data.DeleteMovies(id)
}

func (h *Handler) UpdateMoviesHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/movies/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var updatemovies struct {
		Title       string `json:"title"`
		Genre       string `json:"genre"`
		Year        int    `json:"year"`
		Description string `json:"description"`
	}
	json.NewDecoder(r.Body).Decode(&updatemovies)
	h.data.UpdateMoviesByID(id, updatemovies)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatemovies)
}
