package publisher

import model "books-api/internal/models"

func (p *PublisherManager) GetAll() ([]model.Publisher, error) {
	rows, err := p.DB.Query(`SELECT id, name, created_at FROM publishers ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publishers []model.Publisher
	for rows.Next() {
		var publisher model.Publisher
		err := rows.Scan(&publisher.ID, &publisher.Name, &publisher.CreatedAt)
		if err != nil {
			return nil, err
		}
		publishers = append(publishers, publisher)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return publishers, nil
}
