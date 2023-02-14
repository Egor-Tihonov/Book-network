// Package model ...
package models

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
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Email    string `json:"email"`
	JoinDate string `json:"joinDate"`
}

type UserUpdate struct {
	Status      string `json:"status"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	Name        string `json:"name"`
	Username    string `json:"username"`
}

type LastUsers struct {
	Id       string `json:"Id"`
	Username string `json:"username"`
}

type GetOtherUserResponse struct {
	User         *User
	Subscription bool `json:"subscription"`
}
