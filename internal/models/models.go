package models

//{title: "12342", genre: "2345", year: 2345, description: "2345"}

type Movies struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
