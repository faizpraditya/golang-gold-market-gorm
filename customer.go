package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Kalau di db sudah ada table. Kemudian kita modify fieldnya, maka field lama yang sudah hapus (dimodel terbaru) tetap ada di database
// PascalCase to snake_case
// env buat develop
// Production debugnya harus dicopot.
// better db.conn dibungkus dengan struct kita sendiri
type Customer struct {
	gorm.Model
	ID        string `gorm:"column:id;size:36;primaryKey"`
	FirstName string `gorm:"column:first_name;size:50;not null"`
	LastName  string `gorm:"column:last_name;size:50;not null"`
	BirthDate time.Time
	// Address   string
	Addresses []Address
	Status    int
	// Setelah membuat user credential (one to one)
	// Bisa ga pakai struct tag, karena struct foregin key Customer, dan di UserCredential CustomerID
	UserCredential UserCredential    `gorm:"foreignKey:CustomerID"`
	Products       []CustomerProduct `gorm:"many2many:customer_has_products"`
}

func (c *Customer) ToString() string {
	customer, err := json.Marshal(c)
	if err != nil {
		return ""
	}

	return string(customer)
}

func (c *Customer) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("Before save")
	return
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create")
	c.ID = uuid.NewString()

	return
}

func (c *Customer) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("After save")
	return
}

func (c *Customer) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create")
	return
}

// ● gorm.Model -> mencakup field ID, CreatedAt, UpdatedAt, DeletedAt
// ● Customers : Untuk konfigurasi field tags, bisa dilihat di https://gorm.io/docs/models.html#Fields-Tags
// ● Auto Migration Table
// db.AutoMigrate(&Customer)
