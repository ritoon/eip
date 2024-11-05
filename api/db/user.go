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

func DeleteUser(uuidUser string) error {
	if _, err := GetUser(uuidUser); err != nil {
		return fmt.Errorf("db: deleteUser %w", err)
	}
	delete(us, uuidUser)
	return nil

}
