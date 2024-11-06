package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	DBField
	UUIDOwner string `json:"uuid_owner" gorm:"index;colum:uuid_owner"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Lat       string `json:"lat"`
	Lng       string `json:"lng"`
}

func (Address) TableName() string {
	return "addresses"
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.UUID = uuid.New().String()

	return
}
