package publisher

import (
	model "books-api/internal/models"
	"database/sql"
	"fmt"
)

func (p *PublisherManager) GetOne(id int) (*model.Publisher, error) {
	var publisher model.Publisher

	err := p.DB.QueryRow(`
		SELECT id, name, created_at
		FROM publisher
		WHERE id = $1
	`, id).Scan(&publisher.ID, &publisher.Name, &publisher.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("издательство с %d не найдено", id)
	}

	if err != nil {
		return nil, err
	}

	return &publisher, nil
}
