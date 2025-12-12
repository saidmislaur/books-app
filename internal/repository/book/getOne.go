package book

import (
	models "books-api/internal/models"
	"database/sql"
	"fmt"
)

func (br *BookRepo) GetOne(id int) (*models.Book, error) {
	var book models.Book

	err := br.DB.QueryRow(`
		SELECT id, name, description, author_id, category_id, publisher_id, image, file, created_at
		FROM books
		WHERE ID = $1
	`, id).Scan(&book.ID, &book.Name,
		&book.Description, &book.CategoryID,
		&book.PublisherID, &book.Image,
		&book.File, &book.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("книга с id %d – не найдена", id)
	}

	if err != nil {
		return nil, err
	}

	return &book, nil
}
