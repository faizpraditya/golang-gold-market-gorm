package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	Db *gorm.DB
}

func NewDbConn() *DBConn {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	env := "dev"

	// Open balikannya *gorm.DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if env == "dev" {
		return &DBConn{
			Db: db.Debug(),
		}
	}

	return &DBConn{
		Db: db,
	}
}

func (d *DBConn) Close() {
	db, err := d.Db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Close()

	if err != nil {
		log.Fatalln(err)
	}
}

func (d *DBConn) Migration(tables ...interface{}) {
	err := d.Db.AutoMigrate(tables...)
	if err != nil {
		log.Fatalln(err)
	}
}
