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
		return fmt.Errorf("db: createUser %w", err)
	}
	u.UUID = uuid.String()
	us[u.UUID] = u
	return nil
}

func GetUser(uuidUser string) (*model.User, error) {
	//
	return nil, nil
}

func DeleteUser(uuidUser string) error {
	//
	return nil
}
