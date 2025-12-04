package author

import (
	"books-api/internal/models"
	authorRepo "books-api/internal/repository/author"
)

type Service struct {
	Repo *authorRepo.AuthorManager
}

func New(repo *authorRepo.AuthorManager) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetAll() ([]models.Author, error) {
	return s.Repo.GetAll()
}

func (s *Service) GetOne(id int) (*models.Author, error) {
	return s.Repo.GetOne(id)
}

func (s *Service) Create(a models.Author) (*models.Author, error) {
	auth, err := s.Repo.Create(a)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (s *Service) Update(id int, a models.Author) (*models.Author, error) {
	auth, err := s.Repo.Update(id, a)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (s *Service) Delete(id int) error {
	return s.Repo.Delete(id)
}
