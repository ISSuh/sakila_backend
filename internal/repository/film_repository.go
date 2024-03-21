package repository

import (
	"fmt"

	"github.com/ISSuh/monolith-sakila/internal/common"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/model"
	"github.com/ISSuh/monolith-sakila/pkg/db"
)

type FilmRepository interface {
	FilmById(filmId int) (*model.Film, error)
	Films(page *common.Pagnation) (*common.Pagnation, error)
}

type filmRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewFilmRepository(l logger.Logger, d *db.Database) FilmRepository {
	return &filmRepository{
		log: l,
		db:  d,
	}
}

func (r *filmRepository) FilmById(filmId int) (*model.Film, error) {
	e := r.db.Engine()

	film := new(model.Film)
	if err := e.Model(&film).Preload("Language").Preload("OriginalLanguage").Find(&film, filmId).Error; err != nil {
		return nil, err
	}
	return film, nil
}

func (r *filmRepository) Films(page *common.Pagnation) (*common.Pagnation, error) {
	e := r.db.Engine()

	total, err := r.TotalFilms()
	if err != nil {
		return nil, err
	}

	films := []*model.Film{}
	if err := e.Scopes(db.Pagnation(e, page)).Find(&films).Error; err != nil {
		return nil, err
	}

	fmt.Printf("page : %d, offset : %d, limit : %d, total : %d, films len : %d\n",
		page.Page, page.Offset, page.Limit, *total, len(films))

	page.Total = *total
	page.Item = films
	return page, nil
}

func (r *filmRepository) TotalFilms() (*int64, error) {
	e := r.db.Engine()

	total := new(int64)
	if err := e.Model(&model.Film{}).Count(total).Error; err != nil {
		return nil, err
	}
	return total, nil
}
