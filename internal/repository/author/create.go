package author

import (
	model "books-api/internal/models"
	"errors"
	"fmt"
)

func (am *AuthorManager) Create(a model.Author) (model.Author, error) {
	if a.Name == "" {
		return model.Author{}, errors.New("имя автора обязательно")
	}

	var author model.Author
	err := am.DB.QueryRow(`
		INSERT INTO authors (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at
	`, a.Name, a.Description).Scan(
		&author.ID, &author.Name, &author.Description, &author.CreatedAt,
	)

	if err != nil {
		return model.Author{}, fmt.Errorf("ошибка создания автора: %w", err)
	}

	return author, nil
}
