package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// "github.com/eroub/journal-api-go/internal/app/model"
)

var DB *gorm.DB

func SetupDatabaseConnection() {
    var err error

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"), 
        os.Getenv("DB_PASS"), 
        os.Getenv("DB_SERVER"), 
        "trade_journal") // Database name

    // Enable detailed logging
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second, // Slow SQL threshold
            LogLevel:      logger.Info, // Log level
            Colorful:      true,        // Enable color
        },
    )

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // AutoMigrate
    // DB.AutoMigrate(&model.User{}, &model.Account{}, &model.Transaction{}, &model.Journal{})
}

func CloseDatabaseConnection() {
    db, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to close database connection: %v", err)
    }

    db.Close()
}
