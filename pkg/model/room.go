package model

type Room struct {
	Id    string `json:"id"`
	Name  string `json:"login"`
	Owner User   `json:"owner"`
	Users []User `json:"users"`
}
