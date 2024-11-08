package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Address is a struct that represents an address.
type Address struct {
	DBField
	// UUIDOwner is the UUID of the owner of the address.
	UUIDOwner string `json:"uuid_owner" gorm:"index;colum:uuid_owner"`
	// Street is the street of the address ex: 123 Main St.
	Street string `json:"street"`
	// City is the city of the address ex: New York.
	City string `json:"city"`
	// State is the state of the address ex: NY.
	State string `json:"state"`
	// Zip is the zip code of the address ex: 10001.
	Zip string `json:"zip"`
	// Lat is the latitude of the address.
	Lat float64 `json:"lat"`
	// Lng is the longitude of the address.
	Lng float64 `json:"lng"`
}

// TableName is a method that implements the gorm.Tabler interface.
// It returns the table name for the address.
func (Address) TableName() string {
	return "addresses"
}

// BeforeCreate is a method that implements the gorm.BeforeCreateInterface.
// It generates a new UUID for the address when you want to create a new address.
func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.UUID = uuid.New().String()
	return
}

func (a *Address) String() string {
	return a.Street + ", " + a.City + ", " + a.State + " " + a.Zip
}
