package db

import (
	"context"
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

func (db *DB) UpdateImage(uuidGame, uriImage string) (*model.Game, error) {
	err := db.dbConn.Model(&model.Game{}).Where("uuid = ?", uuidGame).Updates(map[string]interface{}{
		"uri_image": uriImage,
	}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NewErrorNotFound("getGame", fmt.Errorf("db: getGame %q not found", uuidGame))
		}
		return nil, NewErrorInternal("getGame", err)
	}

	return db.GetGame(uuidGame)
}

func (db *DB) DeleteGame(uuidGame string) *Error {
	if _, err := db.GetGame(uuidGame); err != nil {
		return &Error{Err: err, Message: "deleteGame", Code: 404}
	}
	db.dbConn.Where("uuid = ?", uuidGame).Delete(&model.Game{})
	return nil
}

func (db *DB) SearchGames(ctx context.Context, name string) ([]model.Game, error) {
	var games []model.Game
	if name != "" {
		err := db.dbConn.Where("name = ?", name).Find(&games).Error
		if err != nil {
			return nil, NewErrorInternal("searchGames", err)
		}
		return games, nil
	}

	err := db.dbConn.Find(&games).Error
	if err != nil {
		return nil, NewErrorInternal("searchGames", err)
	}
	return games, nil
}
