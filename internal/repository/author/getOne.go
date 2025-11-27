package author

import (
	model "books-api/internal/models"
	"database/sql"
	"fmt"
)

func (am *AuthorManager) GetOne(id int) (*model.Author, error) {
	var author model.Author

	err := am.DB.QueryRow(`
		SELECT id, name, description, created_at
		FROM authors
		WHERE id = $1
	`, id).Scan(&author.ID, &author.Name, &author.Description, &author.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("автор с id %d не найден", id)
	}

	if err != nil {
		return nil, err
	}

	return &author, nil
}
