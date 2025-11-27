package author

import (
	model "books-api/internal/models"
	"database/sql"
	"fmt"
)

func (am *AuthorManager) Update(id int, updated model.Author) (model.Author, error) {
	var existing model.Author
	err := am.DB.QueryRow(`
		SELECT id, name, description, created_at
		FROM authors
		WHERE id = $1
	`, id).Scan(&existing.ID, &existing.Name, &existing.Description, &existing.CreatedAt)

	if err == sql.ErrNoRows {
		return model.Author{}, fmt.Errorf("автор с id %d не найден", id)
	}

	if err != nil {
		return model.Author{}, err
	}

	if updated.Name != "" {
		existing.Name = updated.Name
	}

	if updated.Description != "" {
		existing.Description = updated.Description
	}

	err = am.DB.QueryRow(`
		UPDATE authors
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id, name, description, created_at
	`, existing.Name, existing.Description, id).
		Scan(&existing.ID, &existing.Name, &existing.Description, &existing.CreatedAt)

	if err != nil {
		return model.Author{}, err
	}

	return existing, nil
}
