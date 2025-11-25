package category

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

// CREATE TABLE categories (
//     id SERIAL PRIMARY KEY,
//     name VARCHAR(255) NOT NULL,
//     image TEXT,
//     created_at TIMESTAMP NOT NULL DEFAULT NOW()
// );
