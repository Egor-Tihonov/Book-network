package models

type Feed struct {
	Post
	Status     string `json:"status"`
	Username   string `json:"username"`
	CreateDate string `json:"createDate"`
	UserId     string `json:"userid"`
}

type Post struct {
	PostId        string `json:"postId"`
	AuthorName    string `json:"authorName"`
	AuthorSurname string `json:"authorSurname"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}
