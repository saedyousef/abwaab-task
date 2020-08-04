package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// Initilaize the connection to sqlite database.
func ConnectDatabase() {
    database, err := gorm.Open("sqlite3", "test.db")

    if err != nil {
        panic("Failed to connect to database!")
    }

    // Migrating models.
    database.AutoMigrate(&User{})
    database.AutoMigrate(&Tweet{})

    DB = database
}