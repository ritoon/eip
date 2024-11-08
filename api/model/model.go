package model

import "time"

// DBField is a struct that contains the common fields for all the models.
type DBField struct {
	// UUID is the UUID of the model.
	UUID string `json:"uuid" gorm:"primaryKey"`
	// CreatedAt is the time when the model was created.
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	// UpdatedAt is the time when the model was updated.
	DeletedAt time.Time `json:"deleted_at"`
}
