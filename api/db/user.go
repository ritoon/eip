package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/ritoon/eip/api/model"
)

func (db *DB) CreateUser(u *model.User) error {
	// Create a user.
	log.Println("CreateUser u: ", u)
	db.dbConn.Create(u)

	// u2 := &model.User{
	// 	Name:    "jinzhu",
	// 	Address: model.Address{City: "Paris"},
	// }
	// db.dbConn.Create(u2)

	return nil
}

func (db *DB) GetUser(uuidUser string) (*model.User, error) {
	u := model.User{}
	err := db.dbConn.Model(&model.User{}).Preload("Games").Where("uuid = ?", uuidUser).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NewErrorNotFound("getUser", fmt.Errorf("db: getUser %q not found", uuidUser))
		}
		return nil, NewErrorInternal("getUser", err)
	}

	return &u, nil
}

func (db *DB) DeleteUser(uuidUser string) *Error {
	if _, err := db.GetUser(uuidUser); err != nil {
		return &Error{Err: err, Message: "deleteUser", Code: 404}
	}
	db.dbConn.Where("uuid = ?", uuidUser).Delete(&model.User{})
	return nil
}

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
