package author

import (
	model "books-api/internal/models"
)

func (am *AuthorManager) GetAll() ([]model.Author, error) {
	rows, err := am.DB.Query(`SELECT id, name, description, created_at FROM authors ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []model.Author
	for rows.Next() {
		var author model.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Description, &author.CreatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}
