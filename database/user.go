package database

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}
