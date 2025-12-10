package publisher

import (
	model "books-api/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

func (p *PublisherManager) Update(id int, updated model.Publisher) (model.Publisher, error) {
	var existing model.Publisher
	err := p.DB.QueryRow(`
		SELECT id, name, created_at
		FROM publishers
		WHERE id = $1
	`, id).Scan(&existing.ID, &existing.Name, &existing.CreatedAt)

	if err == sql.ErrNoRows {
		return model.Publisher{}, fmt.Errorf("издательства с id %d не найдено", id)
	}

	if err != nil {
		return model.Publisher{}, err
	}

	if updated.Name != "" {
		existing.Name = updated.Name
	}

	err = p.DB.QueryRow(`
		UPDATE publishers
		SET name = $1
		WHERE id = $2
		RETURNING id, name, created_at
	`, existing.Name, existing.ID).Scan(&existing.ID, &existing.Name, &existing.CreatedAt)

	if err != nil {
		return model.Publisher{}, errors.New("ошибка при обновлении издательства")
	}

	return existing, nil
}
