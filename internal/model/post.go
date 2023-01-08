package model

type Post struct {
	UserID        string `json:"userid"`
	AuthorName    string `json:"authorName"`
	AuthorSurname string `json:"authorSurname"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}
type Author struct {
	ID            string `json:"id"`
	AuthorName    string `json:"authorName"`
	AuthorSurname string `json:"authorSurname"`
}
