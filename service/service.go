package service

import (
	"roketin/entity"
	"roketin/input"
	"roketin/repository"
)

type Service interface {
	Create(input input.Film) (entity.Film, error)
}
type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) Create(input input.Film) (entity.Film, error) {
	film := entity.Film{}

	film.Title = input.Title
	film.Description = input.Description
	film.Artist = input.Artist
	film.Duration = input.Duration
	film.Filename = input.Filename
	film.Genre = input.Genre

	newFilm, err := s.repository.Create(film)

	if err != nil {
		return film, err
	}
	return newFilm, nil
}

// func (s *service) Update(ID int, input input.Film) (entity.Film, error) {
// 	nakes, err := s.repository.FindIDNakes(ID)
// 	if err != nil {
// 		return nakes, err
// 	}
// 	nakes.FaskesID = input.FaskesID
// 	nakes.JumlahNakes = input.JumlahNakes
// 	nakes.Addres = input.Addres
// 	nakes.Name = input.Name

// 	newNakes, err := s.repository.UpdateNakes(nakes)

// 	if err != nil {
// 		return nakes, err
// 	}
// 	return newNakes, nil
// }
// func (s *service) FindIDNakes(ID int) (entity.Nakes, error) {
// 	nakes, err := s.repository.FindIDNakes(ID)
// 	if err != nil {
// 		return nakes, err
// 	}
// 	if nakes.ID == 0 {
// 		return nakes, errors.New("nakes not found")
// 	}
// 	return nakes, nil
// }
// func (s *service) FindAllNakes() ([]entity.Nakes, error) {
// 	nakes, err := s.repository.FindAllNakes()
// 	if err != nil {
// 		return nakes, err
// 	}
// 	return nakes, err
// }
// func (s *service) DeleteNakes(ID int) (entity.Nakes, error) {
// 	nakes, err := s.repository.DeleteNakes(ID)
// 	if err != nil {
// 		return nakes, err
// 	}
// 	return nakes, err
// }
