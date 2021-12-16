package main

import "gorm.io/gorm"

type BaseRepo struct {
	conn *DBConn
}

func (br *BaseRepo) HandleError(db *gorm.DB) error {
	if db.Error != nil {
		return db.Error
	}
	return nil
}
