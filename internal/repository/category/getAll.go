package repository

import models "books-api/internal/models"

func (cm *Repository) GetCategories() ([]models.Category, error) {
	rows, err := cm.DB.Query(`SELECT id, name, image, created_at FROM categories ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Image, &category.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
