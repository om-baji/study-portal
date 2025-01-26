package initialisers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var clientDb *gorm.DB

func ConnectDb() *gorm.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Database connection Failed!")
	}

	url := os.Getenv("DATABASE_URL")
	clientDb, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection Failed!")
	}

	fmt.Println("Database Connected!")

	return clientDb
}
