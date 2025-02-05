package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Инициализация БД
func Init() error {
    db, err := gorm.Open(sqlite.Open("./workouts.db"), &gorm.Config{})
    if err != nil {
        return err
    }
    
    DB = db
    return nil
}

// Автомиграция
func AutoMigrate() error {
    return DB.AutoMigrate(
        &User{},
        &Schedule{},
        &Reminder{},
    )
}

