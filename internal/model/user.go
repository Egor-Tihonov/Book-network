// Package model ...
package model

// UserModel for create user
type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `bson,json:"password"`
	Email    string `json:"email"`
}

//User for response to ckient
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Email    string `json:"email"`
	JoinDate string `json:"joinDate"`
}

type UserUpdate struct {
	Status   string `json:"status"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
