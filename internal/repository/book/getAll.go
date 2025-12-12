package book

import models "books-api/internal/models"

func (br *BookRepo) GetAll() ([]models.Book, error) {
	rows, err := br.DB.Query(`
		SELECT id, name, description, author_id, category_id, publisher_id, image, file, created_at 
		FROM books
		ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID, &book.Name,
			&book.Description, &book.AuthorID,
			&book.CategoryID, &book.PublisherID,
			&book.Image, &book.File, &book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
