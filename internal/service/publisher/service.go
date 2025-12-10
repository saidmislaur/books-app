package publisher

import (
	"books-api/internal/models"
	publishRepi "books-api/internal/repository/publisher"
)

type Service struct {
	Repo *publishRepi.PublisherManager
}

func New(repo *publishRepi.PublisherManager) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetAll() ([]models.Publisher, error) {
	return s.Repo.GetAll()
}

func (s *Service) GetOne(id int) (*models.Publisher, error) {
	return s.Repo.GetOne(id)
}

func (s *Service) Create(p models.Publisher) (*models.Publisher, error) {
	publisher, err := s.Repo.Create(p)
	if err != nil {
		return nil, err
	}

	return &publisher, nil
}

func (s *Service) Update(id int, p models.Publisher) (*models.Publisher, error) {
	publisher, err := s.Repo.Update(id, p)
	if err != nil {
		return nil, err
	}

	return &publisher, nil
}

func (s *Service) Delete(id int) error {
	return s.Repo.Delete(id)
}
