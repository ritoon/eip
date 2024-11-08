package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is a struct that represents a user.
type User struct {
	DBField
	// Name is the name of the user.
	Name string `json:"name"`
	// Address is the address of the user.
	Address Address `json:"address,omitempty" gorm:"foreignKey:UUIDOwner"`
	// Games is the games owned by this user.
	Games []Game `json:"games,omitempty" gorm:"foreignKey:UUIDOwner"`
	UserLogin
}

// TableName is a method that implements the gorm.Tabler interface.
// It returns the table name for the user.
func (User) TableName() string {
	return "users"
}

// BeforeCreate is a method that implements the gorm.BeforeCreateInterface.
// It generates a new UUID for the user when you want to create a new user.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	return
}

// UserLogin is a struct that represents the user login.
// It contains the email and password.
type UserLogin struct {
	// Email is the email of the user.
	Email string `json:"email" form:"email"`
	// Pass is the password of the user in hash with SHA256.
	Pass *Password `json:"pass,omitempty" form:"pass"`
}

// ValidateLogin is a method that validates the user login.
// It returns an error if the email or password is empty.
func (us *UserLogin) ValidateLogin() error {
	if us.Email == "" {
		return fmt.Errorf("email is required")
	}
	if us.Pass == nil {
		return fmt.Errorf("password is required")
	}
	return nil
}

// Password is a type that represents the password.
type Password string

// UnmarshalJSON is a method that implements the json.Unmarshaler interface.
// It hashes the password using SHA256 and stores it in the Password type.
// By doing this, the password is never stored in plain text.
// - Be aware that if you are using another method for binding with gin, you need to implement the same logic.
// - Be aware that when using this method, the password is stored in the database as a hash.
func (p *Password) UnmarshalJSON(b []byte) error {
	// use an auxiliary string to unmarshal the password
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	// and then hash it using SHA256
	sum := sha256.Sum256([]byte(s))
	val := fmt.Sprintf("%x", sum)
	// store the hashed password in the Password type
	*p = Password(val)
	return nil
}

// MarshalJSON is a method that implements the json.Marshaler interface.
// It returns a string with the value "********".
// This is used to avoid sending the password in the response.
func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal("********")
}
