package models

type User struct {
	Id       int64
	Username string
	Name     string
	Phone    string
	IsAdmin  bool
}
