package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AuthorID    int       `json:"author_id"`
	CategoryID  int       `json:"category_id"`
	PublisherID int       `json:"publisher_id"`
	Image       string    `json:"image"`
	File        string    `json:"file"`
	CreatedAt   time.Time `json:"created_at"`
}
