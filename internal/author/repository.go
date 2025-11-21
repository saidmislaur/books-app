package author

import (
	"database/sql"
	"errors"
	"fmt"
)

func NewManager(db *sql.DB) *AuthorManager {
	return &AuthorManager{DB: db}
}

func (am *AuthorManager) GetAuthors() ([]Author, error) {
	rows, err := am.DB.Query("SELECT id, name, description, created_at FROM authors ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		var author Author
		err := rows.Scan(&author.ID, &author.Name, &author.Description, &author.CreatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (am *AuthorManager) GetAuthorById(id int) (*Author, error) {
	var author Author

	err := am.DB.QueryRow(`
		SELECT id, name, description, created_at
		FROM authors
		WHERE id = $1
	`, id).Scan(&author.ID, &author.Name, &author.Description, &author.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("автор с id %d не найден", id)
	}

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (am *AuthorManager) CreateAuthor(a Author) (Author, error) {
	if a.Name == "" {
		return Author{}, errors.New("имя автора обязательно")
	}

	var author Author
	err := am.DB.QueryRow(`
		INSERT INTO authors (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at
	`, a.Name, a.Description).Scan(
		&author.ID, &author.Name, &author.Description, &author.CreatedAt,
	)

	if err != nil {
		return Author{}, fmt.Errorf("ошибка создания автора: %w", err)
	}

	return author, nil
}

func (am *AuthorManager) UpdateAuthor(id int, updated Author) (Author, error) {
	var existing Author
	err := am.DB.QueryRow(`
		SELECT id, name, description, created_at
		FROM authors
		WHERE id = $1
	`, id).Scan(&existing.ID, &existing.Name, &existing.Description, &existing.CreatedAt)

	if err == sql.ErrNoRows {
		return Author{}, fmt.Errorf("автор с id %d не найден", id)
	}

	if err != nil {
		return Author{}, err
	}

	name := existing.Name
	if updated.Name != "" {
		name = updated.Name
	}
	description := existing.Description
	if updated.Description != "" {
		description = updated.Description
	}

	err = am.DB.QueryRow(`
		UPDATE authors
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id, name, description, created_at
	`, name, description, id).Scan(&existing.ID, &existing.Name, &existing.Description, &existing.CreatedAt)

	if err == sql.ErrNoRows {
		return Author{}, fmt.Errorf("автор с id %d не найден", id)
	}
	if err != nil {
		return Author{}, err
	}

	return existing, nil
}

func (am *AuthorManager) DeleteAuthor(id int) error {
	res, err := am.DB.Exec(`DELETE FROM authors WHERE id = $1`, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("автор с id %d не найден", id)
	}

	return nil
}
