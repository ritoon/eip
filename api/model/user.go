package model

// User is a struct that represents a user.
type User struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
