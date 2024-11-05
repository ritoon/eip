package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ritoon/eip/api/model"
)

var us = map[string]*model.User{}

func CreateUser(u *model.User) error {
	// Create a user.
	uuid, err := uuid.NewRandom()
	if err != nil {
		return NewErrorInternal("createUser", err)
	}
	u.UUID = uuid.String()
	us[u.UUID] = u
	return nil
}

func GetUser(uuidUser string) (*model.User, error) {
	u, ok := us[uuidUser]

	if !ok {
		return nil, NewErrorNotFound("getUser", fmt.Errorf("db: getUser %q not found", uuidUser))
	}
	return u, nil
}

func DeleteUser(uuidUser string) *Error {
	if _, err := GetUser(uuidUser); err != nil {
		return &Error{Err: err, Message: "deleteUser", Code: 404}
	}
	delete(us, uuidUser)
	return nil

}
