package domain

type Game struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
