package repository

import (
	"fmt"
)

func (cm *Repository) DeleteCategory(id int) error {
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
