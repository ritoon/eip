package model

import (
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

// Game is a struct that represents a game.
type Game struct {
	DBField
	// UUIDOwner is the UUID of the owner of the game.
	UUIDOwner string `json:"uuid_owner" gorm:"index"`
	// Name is the name of the game.
	Name string `json:"name"`
	// Age for playing with the game.
	Age string `json:"age"`
	// Type of the game.
	Type string `json:"type"`
	// PlayersNb is the number of players for the game.
	PlayersNb string `json:"players_nb"`
	// URIImage is the URI of the image of the game.
	URIImage string `json:"uri_image"`
}

// TableName is a method that implements the gorm.Tabler interface.
// It returns the table name for the game.
func (Game) TableName() string {
	return "games"
}

// BeforeCreate is a method that implements the gorm.BeforeCreateInterface.
// It generates a new UUID for the game when you want to create a new game.
func (g *Game) BeforeCreate(tx *gorm.DB) (err error) {
	g.UUID = "gam-" + ksuid.New().String()
	return
}
