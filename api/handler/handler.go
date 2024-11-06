package handler

import (
	"github.com/ritoon/eip/api/cache"
	"github.com/ritoon/eip/api/db"
)

type Handler struct {
	cache *cache.Redis
	db    *db.DB
}

func New(cache *cache.Redis, db *db.DB) *Handler {
	return &Handler{cache: cache, db: db}
}
