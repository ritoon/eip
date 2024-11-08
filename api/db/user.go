package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/ritoon/eip/api/model"
)

// CreateUser create user in db and return error if any.
func (db *DB) CreateUser(u *model.User) error {
	// Create a user.
	log.Println("CreateUser u: ", u)
	db.dbConn.Create(u)
	return nil
}

// GetUser get user from db and return error if any.
func (db *DB) GetUser(uuidUser string) (*model.User, error) {
	// create a user.
	u := model.User{}
	// Preload("Games") is used to load the games of the user.
	// Preload("Address") is used to load the address of the user.
	err := db.dbConn.Model(&model.User{}).Preload("Games").Preload("Address").Where("uuid = ?", uuidUser).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NewErrorNotFound("getUser", fmt.Errorf("db: getUser %q not found", uuidUser))
		}
		return nil, NewErrorInternal("getUser", err)
	}

	return &u, nil
}

// UpdateUser update user in db and return error if any.
func (db *DB) DeleteUser(uuidUser string) *Error {
	if _, err := db.GetUser(uuidUser); err != nil {
		return &Error{Err: err, Message: "deleteUser", Code: 404}
	}
	db.dbConn.Where("uuid = ?", uuidUser).Delete(&model.User{})
	return nil
}

// GetUserByEmail get user from db and return error if any.
func (db *DB) GetUserByEmail(email string) (*model.User, error) {
	u := model.User{}
	err := db.dbConn.Model(&model.User{}).Where("email = ?", email).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			NewErrorNotFound("getUserByEmail", fmt.Errorf("db: getUserByEmail %q not found", email))
		}
		return nil, NewErrorInternal("getUserByEmail", err)
	}
	return &u, nil
}
