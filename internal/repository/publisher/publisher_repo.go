package publisher

import "database/sql"

type PublisherManager struct {
	DB *sql.DB
}

func NewManager(db *sql.DB) *PublisherManager {
	return &PublisherManager{DB: db}
}
