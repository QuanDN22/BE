package models

import (
	"log"

	"gorm.io/gorm"
)

func Init_DB(db *gorm.DB) {
	var err error

	// delete table if it doesn't exist
	err = db.Migrator().DropTable(&Server{})
	if err != nil {
		log.Fatalf("Failed to drop table servers: %v", err)
	} else {
		log.Println("Dropped table servers")
	}

	err = db.Migrator().DropTable(&ServerDeleted{})
	if err != nil {
		log.Fatalf("Failed to drop table server_deleteds: %v", err)
	} else {
		log.Println("Dropped table server_deleted")
	}

	// Auto migrate the Server model
	err = db.AutoMigrate(&Server{})
	if err != nil {
		log.Fatalf("Failed to migrate servers datable: %v", err)
	} else {
		log.Println("migrate servers datable successfully")
	}

	err = db.AutoMigrate(&ServerDeleted{})
	if err != nil {
		log.Fatalf("Failed to migrate server_deleteds datable: %v", err)
	} else {
		log.Println("migrate server_deleted successfully")
	}
}
