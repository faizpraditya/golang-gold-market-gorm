package main

import "gorm.io/gorm"

type CustomerProduct struct {
	ID          string     `gorm:"column:id;size:36;primaryKey"`
	ProductName string     `gorm:"column:product_name;size:50;not null"`
	Customers   []Customer `gorm:"many2many:customer_has_products"`
	gorm.Model
}
