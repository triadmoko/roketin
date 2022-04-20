package repository

import (
	"roketin/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Create(film entity.Film) (entity.Film, error)
	UpdateFilm(film entity.Film) (entity.Film, error)
	FindID(ID int) (entity.Film, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(film entity.Film) (entity.Film, error) {
	err := r.db.Create(&film).Error
	if err != nil {
		return film, err
	}
	return film, nil
}

func (r *repository) UpdateFilm(film entity.Film) (entity.Film, error) {
	err := r.db.Save(&film).Error
	if err != nil {
		return film, err
	}
	return film, nil
}

func (r *repository) ListFilm() ([]entity.Film, error) {
	var film []entity.Film
	err := r.db.Find(&film).Error
	if err != nil {
		return film, err
	}
	return film, nil
}

func (r *repository) FindID(ID int) (entity.Film, error) {
	var film entity.Film
	err := r.db.Where("ID = ?", ID).Find(&film).Error
	if err != nil {
		return film, err
	}
	return film, nil
}
