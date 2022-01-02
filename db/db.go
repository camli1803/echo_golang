package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("todolist1.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}
