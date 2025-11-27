package models

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryManager struct {
	DB *sql.DB
}
