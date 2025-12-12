package book

import (
	models "books-api/internal/models"
	"fmt"
)

func (br *BookRepo) Create(b models.Book) (models.Book, error) {
	var book models.Book

	err := br.DB.QueryRow(
		`INSERT INTO books (name, description, author_id, category_id, publisher_id, image, file)
		VALUE ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, description, author_id, category_id, publisher_id, image, file, created_at
		`, b.Name, b.Description, b.AuthorID, b.CategoryID, b.PublisherID, b.Image, b.File).Scan(
		&book.ID, &book.Name, &book.Description,
		&book.AuthorID, &book.CategoryID, &book.PublisherID, &book.Image,
		&book.File, &book.CreatedAt)

	if err != nil {
		return models.Book{}, fmt.Errorf("ошибка при добавлении книги: %w", err)
	}

	return book, nil
}
