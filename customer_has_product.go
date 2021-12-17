package main

import "gorm.io/gorm"

type CustomerHasProduct struct {
	CustomerProductID string `gorm:"primaryKey"`
	CustomerID        string `gorm:"primaryKey"`
	gorm.Model
}
