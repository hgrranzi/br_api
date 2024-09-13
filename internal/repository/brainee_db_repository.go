package repository

import (
	"br_api/internal/models/entity"
	"errors"
	"gorm.io/gorm"
)

type BraineeDBRepository struct {
	db *gorm.DB
}

func NewBraineeDBRepository(db *gorm.DB) BraineeRepository {
	return &BraineeDBRepository{db: db}
}

func (r *BraineeDBRepository) Create(brainee *entity.Brainee) error {
	return r.db.Create(brainee).Error
}

func (r *BraineeDBRepository) GetById(id int) (*entity.Brainee, error) {
	var brainee entity.Brainee
	err := r.db.First(&brainee, id).Error
	if err != nil {
		return nil, err
	}
	return &brainee, nil
}

func (r *BraineeDBRepository) GetAll() ([]*entity.Brainee, error) {
	var brainees []*entity.Brainee
	err := r.db.Find(&brainees).Error
	return brainees, err
}

func (r *BraineeDBRepository) FindByTextAndAuthorAndBrand(text, author, brand string) (*entity.Brainee, error) {
	var brainee entity.Brainee
	err := r.db.Where("text = ? AND author = ? AND brand = ?", text, author, brand).First(&brainee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &brainee, nil
}
