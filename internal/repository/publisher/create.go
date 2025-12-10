package publisher

import (
	model "books-api/internal/models"
	"errors"
	"fmt"
)

func (p *PublisherManager) Create(a model.Publisher) (model.Publisher, error) {
	if a.Name == "" {
		return model.Publisher{}, errors.New("название издательства обязательно")
	}

	var publisher model.Publisher
	err := p.DB.QueryRow(`
	 	INSERT INTO publishers (name)
		VALUES ($1)
		RETURNING id, name, created_at
	`, a.Name).Scan(&publisher.ID, &publisher.Name, &publisher.CreatedAt)

	if err != nil {
		return model.Publisher{}, fmt.Errorf("ошибка добавления издетельства %w", err)
	}

	return publisher, nil
}
