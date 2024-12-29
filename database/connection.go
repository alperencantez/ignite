package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alperencantez/ignite/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s dbname=%s password=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
		return
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	err = db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatal("Auto migration failed")
		return
	}

	DB = db
}
