package database

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/Mawarii/sheethappens/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	timezone := os.Getenv("TZ")

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, username, password, dbName, port, timezone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connection to database:", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Character{}, &model.Skill{}, &model.SkillCategory{}, &model.CharacterSkill{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}

func DB() *gorm.DB {
	return db
}
