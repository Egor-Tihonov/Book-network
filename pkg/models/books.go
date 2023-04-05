package models

type Book struct {
	Author
	BookId string `json:"bookid"`
	Title  string `json:"title"`
}

type Author struct {
	AuthorId string `json:"authorid"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}
