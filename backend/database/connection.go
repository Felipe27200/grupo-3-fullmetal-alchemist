package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"alchemy-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	var db *gorm.DB
	var err error

	// Retry logic to wait for MySQL container
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //  enable SQL logging
		})
		if err == nil {
			break
		}
		log.Printf("Waiting for MySQL to be ready... (%d/10)", i)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}

	// Run AutoMigrate for all models
	err = db.AutoMigrate(
		&models.Alchemist{},
		&models.Mission{},
		&models.Material{},
		&models.Transmutation{},
		&models.Audit{},
	)
	if err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
	}

	// Check how many records exist for sanity
	var count int64
	db.Table("alchemists").Count(&count)
	log.Printf("Found %d alchemists in database '%s'\n", count, name)

	DB = db
	log.Println(" Connected to MySQL and migrated tables successfully.")
}
