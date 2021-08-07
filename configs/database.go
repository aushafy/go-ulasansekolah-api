package configs

import (
	"github.com/aushafy/go-ulasansekolah-api/domain/schools"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("ulasansekolah.db"))
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&schools.School{})

	DB = database
}
