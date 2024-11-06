package db

import (
	"fmt"

	"github.com/ritoon/eip/api/model"
	"gorm.io/gorm"
)

func (db *DB) CreateAddress(u *model.Address) error {
	// Create a Address.
	db.dbConn.Create(u)
	return nil
}

func (db *DB) GetAddress(uuidAddress string) (*model.Address, error) {
	u := model.Address{}
	err := db.dbConn.Model(&model.Address{}).Where("uuid = ?", uuidAddress).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NewErrorNotFound("getAddress", fmt.Errorf("db: getAddress %q not found", uuidAddress))
		}
		return nil, NewErrorInternal("getAddress", err)
	}

	return &u, nil
}

func (db *DB) DeleteAddress(uuidAddress string) *Error {
	if _, err := db.GetAddress(uuidAddress); err != nil {
		return &Error{Err: err, Message: "deleteAddress", Code: 404}
	}
	db.dbConn.Where("uuid = ?", uuidAddress).Delete(&model.Address{})
	return nil
}
