package repository

import (
	models "books-api/internal/models"
	"fmt"
)

func (cm *Repository) CreateCategory(c models.Category) (models.Category, error) {
	var category models.Category
	if c.Name == "" {
		return models.Category{}, fmt.Errorf("название категории обязательно")
	}

	err := cm.DB.QueryRow(`
		INSERT INTO categories (name, image)
		VALUES ($1, $2)
		RETURNING id, name, image, created_at
	`, c.Name, c.Image).Scan(&category.ID, &category.Name, &category.Image, &category.CreatedAt)

	if err != nil {
		return models.Category{}, fmt.Errorf("ошибка создания категории: %w", err)
	}

	return category, nil
}
