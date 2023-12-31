package core

import (
	"log"
	"os"

	"github.com/torbatti/karab/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	Db *gorm.DB
}

var DataBase Storage

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database .\n ", err.Error())
		os.Exit(2)
	}
	// db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.Applicant{}, &models.Company{}, &models.Job{})

	DataBase = Storage{Db: db}
}
