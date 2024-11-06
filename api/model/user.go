package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is a struct that represents a user.
type User struct {
	DBField
	Name    string  `json:"name"`
	Address Address `json:"address,omitempty" gorm:"foreignKey:UUIDOwner"`
	Games   []Game  `json:"games,omitempty" gorm:"foreignKey:UUIDOwner"`
	UserLogin
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()

	return
}

type DBField struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Password string

type UserLogin struct {
	Email string    `json:"email"`
	Pass  *Password `json:"pass,omitempty"`
}

func (p *Password) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	sum := sha256.Sum256([]byte(s))
	val := fmt.Sprintf("%x", sum)
	*p = Password(val)

	return nil
}

func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal("********")
}
