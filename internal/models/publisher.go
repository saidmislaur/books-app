package models

import (
	"database/sql"
	"time"
)

type Publisher struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type PublisherManager struct {
	DB *sql.DB
}
