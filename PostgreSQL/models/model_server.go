package models

import (
	"time"

	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	Server_Name   string `json:"server_name" gorm:"column:server_name; not null; unique"`
	Server_IPv4   string `json:"server_ipv4" gorm:"column:server_ipv4; not null; unique"`
	Server_Status bool   `json:"server_status" gorm:"column:server_status; not null"`
}

// store information about server which is deleted
type ServerDeleted struct {
	Server_ID   uint      `json:"server_id" gorm:"column:server_id; not null;"`
	Server_Name string    `json:"server_name" gorm:"column:server_name; not null;"`
	Server_IPv4 string    `json:"server_ipv4" gorm:"column:server_ipv4; not null;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at; not null;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at; not null;"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"column:deleted_at; not null;"`
}
