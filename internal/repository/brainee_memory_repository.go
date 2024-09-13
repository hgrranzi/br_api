package repository

import (
	"br_api/internal/models/entity"
	"errors"
	"sync"
	"time"
)

type BraineeMemoryRepository struct {
	mu       sync.RWMutex
	brainees map[int]*entity.Brainee
	nextId   int
}

func NewBraineeMemoryRepository() BraineeRepository {
	return &BraineeMemoryRepository{
		brainees: make(map[int]*entity.Brainee),
		nextId:   1,
	}
}

func (r *BraineeMemoryRepository) Create(brainee *entity.Brainee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	brainee.Id = r.nextId
	r.nextId++
	brainee.CreatedAt = time.Now()
	r.brainees[brainee.Id] = brainee

	return nil
}

func (r *BraineeMemoryRepository) GetById(id int) (*entity.Brainee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	brainee, exists := r.brainees[id]
	if !exists {
		return nil, errors.New("brainee not found")
	}

	return brainee, nil
}

func (r *BraineeMemoryRepository) GetAll() ([]*entity.Brainee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	brainees := make([]*entity.Brainee, 0, len(r.brainees))
	for _, brainee := range r.brainees {
		brainees = append(brainees, brainee)
	}

	return brainees, nil
}

func (r *BraineeMemoryRepository) FindByTextAndAuthorAndBrand(text, author, brand string) (*entity.Brainee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, brainee := range r.brainees {
		if brainee.Text == text && brainee.Author == author && brainee.Brand == brand {
			return brainee, nil
		}
	}
	return nil, nil
}
