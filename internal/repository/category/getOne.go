package repository

import (
	models "books-api/internal/models"
	"database/sql"
	"fmt"
)

func (cm *Repository) GetCategory(id int) (*models.Category, error) {
	var category models.Category

	err := cm.DB.QueryRow(`
		SELECT id, name, image, created_at
		FROM categories
		WHERE id = $1
	`, id).Scan(&category.ID, &category.Name, &category.Image, &category.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("категория с %d не найдена", id)
	}

	if err != nil {
		return nil, nil
	}

	return &category, nil
}
