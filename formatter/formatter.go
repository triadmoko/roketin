package formatter

import "roketin/entity"

type FormatFilm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`
	Filename    string `json:"filename"`
}

func FormatterFilm(film entity.Film) FormatFilm {
	formatter := FormatFilm{
		Title:       film.Title,
		Description: film.Description,
		Duration:    film.Duration,
		Artist:      film.Artist,
		Genre:       film.Genre,
		Filename:    film.Filename,
	}
	return formatter
}
func FormatFilmAll(faskes []entity.Film) []FormatFilm {
	var filmAll []FormatFilm
	for _, v := range faskes {
		formatter := FormatterFilm(v)
		filmAll = append(filmAll, formatter)
	}
	return filmAll
}
