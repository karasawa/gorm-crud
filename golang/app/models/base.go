package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

  )

type User struct {
	gorm.Model
	Name string
	Email string
}

type Post struct {
	gorm.Model
	Content string
}

func DbInit() *gorm.DB {
	dsn := "host=postgres user=root password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})

	return db
}