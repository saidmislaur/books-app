package book

import "database/sql"

type BookRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) *BookRepo {
	return &BookRepo{DB: db}
}
