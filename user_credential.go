package main

import "gorm.io/gorm"

type UserCredential struct {
	ID         string `gorm:"column:id;size:36;primaryKey"`
	Username   string `gorm:"size:50;not null"`
	Password   string `gorm:"size:10;not null"`
	Email      string `gorm:"size:50;not null"`
	CustomerID string
	IsActive   bool
	gorm.Model
}

// Bisa mengoverride nama tablenya (dari dokumentasi)
func (u *UserCredential) TableName() string {
	return "mst_user_credential"
}
