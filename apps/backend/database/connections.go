package database

import (
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/santiagofelizzola/coaching-app/models"
)

var DB *gorm.DB

func Connect() {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = db
    DB.AutoMigrate(
        &models.User{},
        &models.Team{},
        &models.Task{},
        &models.UserTask{},
        &models.Session{},
    )
    
    log.Println("✅ Connected to database!")
}