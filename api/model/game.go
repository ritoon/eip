package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	DBField
	UUIDOwner string `json:"uuid_owner" gorm:"index"`
	Name      string `json:"name"`
	URIImage  string `json:"uri_image"`
}

func (Game) TableName() string {
	return "games"
}

func (g *Game) BeforeCreate(tx *gorm.DB) (err error) {
	g.UUID = uuid.New().String()

	return
}
