package postgres

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type postgresDB struct {
	db *gorm.DB
}

func NewPostgresDB(
	host string,
	user string,
	password string,
	dbname string,
	port string,
) *postgresDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	} 

	return &postgresDB{
		db: db,
	}
}
