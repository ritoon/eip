package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is a struct that represents a user.
type User struct {
	DBField
	Name string `json:"name"`
	UserLogin
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()

	return
}

type DBField struct {
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

var u = User{
	DBField: DBField{
		UUID: "1234",
	},
	Name: "Ritoon",
	UserLogin: UserLogin{
		Email: "",
		Pass:  "1234",
	},
}

type UserLogin struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
