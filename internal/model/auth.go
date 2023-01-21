package model

type AuthUserModel struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `bson,json:"password"`
}

// AuthenticationForm ...
type AuthenticationForm struct {
	AuthString string `json:"authString"`
	Password   string `json:"password"`
}
