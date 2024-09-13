package service

import (
	"br_api/internal/models/entity"
	"br_api/internal/repository"
	"errors"
)

type BraineeService struct {
	repo repository.BraineeRepository
}

func NewBraineeService(repo repository.BraineeRepository) *BraineeService {
	return &BraineeService{repo: repo}
}

func (s *BraineeService) CreateBrainee(brainee *entity.Brainee) error {
	existing, err := s.repo.FindByTextAndAuthorAndBrand(brainee.Text, brainee.Author, brainee.Brand)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("brainee with the same data already exists")
	}
	return s.repo.Create(brainee)
}

func (s *BraineeService) GetBraineeById(id uint) (*entity.Brainee, error) {
	return s.repo.GetById(id)
}

func (s *BraineeService) GetAllBrainees() ([]*entity.Brainee, error) {
	return s.repo.GetAll()
}
