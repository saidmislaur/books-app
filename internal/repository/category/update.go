package repository

import (
	"database/sql"
	"fmt"

	models "books-api/internal/models"
)

func (cm *Repository) UpdateCategory(id int, c models.Category) (models.Category, error) {
	var existing models.Category
	err := cm.DB.QueryRow(`
		SELECT id, name, image, created_at
		FROM categories
		WHERE id = $1
	`, id).Scan(&existing.ID, &existing.Name, &existing.Image, &existing.CreatedAt)

	if err == sql.ErrNoRows {
		return models.Category{}, fmt.Errorf("категория с %d не найдена", id)
	}

	if err != nil {
		return models.Category{}, err
	}

	if c.Name != "" {
		existing.Name = c.Name
	}

	if c.Image != "" {
		existing.Image = c.Image
	}

	err = cm.DB.QueryRow(`
		UPDATE categories
		SET name = $1, image = $2
		WHERE id = $3
		RETURNING id, name, image, created_at
	`, existing.Name, existing.Image, id).Scan(&existing.ID, &existing.Name, &existing.Image, &existing.CreatedAt)

	if err != nil {
		return models.Category{}, err
	}

	return existing, nil
}
