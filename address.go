package main

type Address struct {
	ID         string `gorm:"column:id;size:36;primaryKey"`
	Streetname string `gorm:"size:50;not null"`
	City       string `gorm:"size:50;not null"`
	PostalCode string `gorm:"size:50;not null"`
	CustomerID string
}

func (a *Address) TableName() string {
	return "mst_user_address"
}
