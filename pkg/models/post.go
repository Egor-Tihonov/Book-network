package models

type Post struct {
	PostId        string `json:"postId"`
	AuthorName    string `json:"authorName"`
	AuthorSurname string `json:"authorSurname"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}

type LastPost struct {
	PostId string `json:"postId"`
	Title  string `json:"title"`
}
