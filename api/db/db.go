package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/ritoon/eip/api/model"
)

type DB struct {
	dbConn *gorm.DB
}

func New() *DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Address{})
	db.AutoMigrate(&model.Game{})
	return &DB{dbConn: db}
}
