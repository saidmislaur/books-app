package category

import (
	"database/sql"
	"fmt"
)

func NewManager(db *sql.DB) *CategoryManager {
	return &CategoryManager{DB: db}
}

func (cm *CategoryManager) GetCategories() ([]Category, error) {
	rows, err := cm.DB.Query(`SELECT id, name, image, created_at FROM categories ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
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

func (cm *CategoryManager) GetCategory(id int) (*Category, error) {
	var category Category

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

func (cm *CategoryManager) CreateCategory(c Category) (Category, error) {
	var category Category
	if c.Name == "" {
		return Category{}, fmt.Errorf("название категории обязательно")
	}

	err := cm.DB.QueryRow(`
		INSERT INTO categories (name, image)
		VALUES ($1, $2)
		RETURNING id, name, image, created_at
	`, c.Name, c.Image).Scan(&category.ID, &category.Name, &category.Image, &category.CreatedAt)

	if err != nil {
		return Category{}, fmt.Errorf("ошибка создания категории: %w", err)
	}

	return category, nil
}

func (cm *CategoryManager) UpdateCategory(id int, c Category) (Category, error) {
	var existing Category
	err := cm.DB.QueryRow(`
		SELECT id, name, image, created_at
		FROM categories
		WHERE id = $1
	`, id).Scan(&existing.ID, &existing.Name, &existing.Image, &existing.CreatedAt)

	if err == sql.ErrNoRows {
		return Category{}, fmt.Errorf("категория с %d не найдена", id)
	}

	if err != nil {
		return Category{}, err
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
		return Category{}, err
	}

	return existing, nil
}

func (cm *CategoryManager) DeleteCategory(id int) error {
	row, err := cm.DB.Exec(`
		DELETE FROM categories
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("категории с таким %d не существует", id)
	}

	return nil
}
