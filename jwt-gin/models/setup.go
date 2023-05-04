package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDataBase() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error Loading . env file")

	// }
	// Dbdriver := os.Getenv("DB_DRIVER")
	// DbHost := os.Getenv("DB_HOST")
	// DbUser := os.Getenv("DB_USER")
	// DbPassword := os.Getenv("DB_PASSWORD")
	// DbName := os.Getenv("DB_NAME")
	// DbPort := os.Getenv("DB_PORT")
	//dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connect error:", err)
	} else {
		fmt.Println("Successfully connect")
	}
	DB.AutoMigrate(&User{})

}
