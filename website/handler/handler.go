package handler

import (
	"github.com/ritoon/eip/api/db"
)

type Handler struct {
	db *db.DB
}

// New create new handler instance with db
func New(db *db.DB) *Handler {
	return &Handler{db: db}
}
