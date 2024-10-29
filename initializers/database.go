package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v sslmode=disable", dbUser, dbPass, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		// TODO: Повыеносить ошибки в енамку
		log.Fatalf("Error: DB is not connected")
	}
}
