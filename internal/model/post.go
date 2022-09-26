package model

type Post struct {
	UserID        string `json:"userid"`
	AuthorName    string `json:"author-name"`
	AuthorSurname string `json:"author-surname"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}
type Author struct {
	ID            string `json:"id"`
	AuthorName    string `json:"author-name"`
	AuthorSurname string `json:"author-surname"`
}
