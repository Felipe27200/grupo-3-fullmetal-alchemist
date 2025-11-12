package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"alchemy-system/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// Retry logic to allow MySQL container to start
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Waiting for MySQL to be ready... (%d/10)", i)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}

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

	DB = db
	log.Println("Connected to MySQL and migrated tables successfully.")
}
