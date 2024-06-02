package models

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect to database
func Connection_DB() *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Loading .env file successfully")
	}

	var varEnv map[string]string
	varEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	} else {
		log.Println("Reading .env file successfully")
	}

	var (
		host     = varEnv["PG_DATABASE_HOST"]
		user     = varEnv["PG_DATABASE_USER"]
		password = varEnv["PG_DATABASE_PASSWORD"]
		dbname   = varEnv["PG_DATABASE_DBNAME"]
		port     = varEnv["PG_DATABASE_PORT"]
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	} else {
		log.Println("database connected")
	}

	return db
}
