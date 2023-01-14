package models

import (
	"fmt"
	"log"

	"github.com/karasawa/gorm-crud.git/config"

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

type Todo struct {
	gorm.Model
	Task string
}

func DbInit() *gorm.DB {
	cfg, _ := config.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", 
			cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Todo{})

	return db
}