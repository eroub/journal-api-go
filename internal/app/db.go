package main

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"github.com/eroub/journal-api-go/internal/app/model"
)

var DB *gorm.DB

func SetupDatabaseConnection() {
    var err error

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"), 
        os.Getenv("DB_PASS"), 
        os.Getenv("DB_SERVER"), 
        "trade_journal") // Database name

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // AutoMigrate goes here if you choose to use it
    DB.AutoMigrate(&model.User{}, &model.Account{}, &model.Transaction{}, &model.TradeJournal{})
}

func CloseDatabaseConnection() {
    db, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to close database connection: %v", err)
    }

    db.Close()
}
