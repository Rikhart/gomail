package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Mail struct {
	gorm.Model
	Subject string `gorm:"column:subject"`
	From    string `gorm:"column:from"`
	To      string `gorm:"column:to"`
}

func Connect() {
	fmt.Println("intentanto conextion")
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Mail{})
}
