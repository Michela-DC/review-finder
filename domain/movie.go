package domain

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
