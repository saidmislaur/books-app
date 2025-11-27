package author

import "database/sql"

type AuthorManager struct {
	DB *sql.DB
}

func NewManager(db *sql.DB) *AuthorManager {
	return &AuthorManager{DB: db}
}
