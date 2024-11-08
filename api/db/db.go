package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/ritoon/eip/api/model"
)

// DB holds the database connection
type DB struct {
	dbConn *gorm.DB
}

// New creates a new database connection
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
