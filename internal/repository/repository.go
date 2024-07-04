package repository

import (
	restapimediasoft "tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft"

	"github.com/jmoiron/sqlx"
)

type FurniruresList interface {
	Create(list restapimediasoft.Furnitures_list) (int, error)
	GetAll() ([]restapimediasoft.Furnitures_list, error)
	GetByID(listID int) (restapimediasoft.Furnitures_list, error)
	DeleteList(listID int) error
	UpdateAll(listID int, input restapimediasoft.UpdateListInput) error
	Update(listID int, input restapimediasoft.UpdateListInput) error
}
type Repository struct {
	FurniruresList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		FurniruresList: NewFurnituresListPostgres(db),
	}
}
