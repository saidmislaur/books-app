package category

import (
	"books-api/internal/models"
	repository "books-api/internal/repository/category"
	"errors"
)

type Service struct {
	Repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetAll() ([]models.Category, error) {
	return s.Repo.GetCategories()
}

func (s *Service) GetOne(id int) (*models.Category, error) {
	return s.Repo.GetCategory(id)
}

func (s *Service) Create(c models.Category) (*models.Category, error) {
	if c.Name == "" {
		return nil, errors.New("название для категории обязательно")
	}
	cat, err := s.Repo.CreateCategory(c)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (s *Service) Update(id int, c models.Category) (*models.Category, error) {
	if c.Name == "" {
		return nil, errors.New("название не может быть пустым")
	}
	cat, err := s.Repo.UpdateCategory(id, c)
	if err != nil {
		return nil, err
	}
	return &cat, err
}

func (s *Service) Delete(id int) error {
	return s.Repo.DeleteCategory(id)
}
