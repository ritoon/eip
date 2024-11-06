package db

import (
	"fmt"

	"github.com/ritoon/eip/api/model"
	"gorm.io/gorm"
)

func (db *DB) CreateGame(u *model.Game) error {
	// Create a Game.
	db.dbConn.Create(u)
	return nil
}

func (db *DB) GetGame(uuidGame string) (*model.Game, error) {
	u := model.Game{}
	err := db.dbConn.Model(&model.Game{}).Where("uuid = ?", uuidGame).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NewErrorNotFound("getGame", fmt.Errorf("db: getGame %q not found", uuidGame))
		}
		return nil, NewErrorInternal("getGame", err)
	}

	return &u, nil
}

func (db *DB) DeleteGame(uuidGame string) *Error {
	if _, err := db.GetGame(uuidGame); err != nil {
		return &Error{Err: err, Message: "deleteGame", Code: 404}
	}
	db.dbConn.Where("uuid = ?", uuidGame).Delete(&model.Game{})
	return nil
}
