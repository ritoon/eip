package db

import (
	"context"
	"fmt"

	"github.com/ritoon/eip/api/model"
	"gorm.io/gorm"
)

// CreateGame create game in db and return error if any.
func (db *DB) CreateGame(u *model.Game) error {
	// Create a Game.
	db.dbConn.Create(u)
	return nil
}

// GetGame get game from db and return error if any.
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

// UpdateGame update game in db and return error if any.
// example: db.UpdateGame("uuid", "name", "new name")
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

// DeleteGame delete game in db and return error if any.
// example: db.DeleteGame("uuid")
func (db *DB) DeleteGame(uuidGame string) *Error {
	if _, err := db.GetGame(uuidGame); err != nil {
		return &Error{Err: err, Message: "deleteGame", Code: 404}
	}
	db.dbConn.Where("uuid = ?", uuidGame).Delete(&model.Game{})
	return nil
}

// SearchGames search games in db and return error if any.
// example: db.SearchGames(ctx, "name")
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

// UpdateGame update game in db and return error if any.
func (db *DB) UpdateGame(uuidGame string, data map[string]interface{}) *Error {
	if _, err := db.GetGame(uuidGame); err != nil {
		return &Error{Err: err, Message: "deleteGame", Code: 404}
	}

	delete(data, "uuid")
	delete(data, "uuid_owner")
	delete(data, "created_at")
	delete(data, "updated_at")
	delete(data, "deleted_at")

	db.dbConn.Model(&model.Game{}).Where("uuid = ?", uuidGame).Updates(data)
	return nil
}
