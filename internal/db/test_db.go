package db

import (
	"hitalentTest/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"
)

func ConnectTest() *gorm.DB {
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "sqlite",
		DSN:        "file::memory:?cache=shared",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Chat{}, &models.Message{}); err != nil {
		panic(err)
	}

	return db
}
