package repository

import "br_api/internal/models/entity"

type BraineeRepository interface {
	Create(brainee *entity.Brainee) error
	GetById(id uint) (*entity.Brainee, error)
	GetAll() ([]*entity.Brainee, error)
	FindByTextAndAuthorAndBrand(text, author, brand string) (*entity.Brainee, error)
}
