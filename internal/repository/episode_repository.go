package repository

import (
	"podcaster/internal/model"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	List() ([]model.Episode, error)
}

type GormEpisodeRepository struct {
	db *gorm.DB
}

func (r *GormEpisodeRepository) List() ([]model.Episode, error) {
	var episodes []model.Episode
	err := r.db.Find(&episodes).Error

	if err != nil {
		return nil, err
	}

	return episodes, nil
}

func NewGormEpisodeRepository(db *gorm.DB) EpisodeRepository {
	return &GormEpisodeRepository{db}
}
