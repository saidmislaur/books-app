package models

import (
	"database/sql"
	"time"
)

type Author struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type AuthorManager struct {
	DB *sql.DB
}
