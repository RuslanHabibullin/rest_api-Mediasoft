package service

import (
	restapimediasoft "tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/repository"
)

type Service struct {
	Furnitures_List
}

type Furnitures_List interface {
	Create(list restapimediasoft.Furnitures_list) (int, error)
	GetAll() ([]restapimediasoft.Furnitures_list, error)
	GetByID(listID int) (restapimediasoft.Furnitures_list, error)
	DeleteList(listID int) error
	UpdateAll(listID int, input restapimediasoft.UpdateListInput) error
	Update(listIDid int, input restapimediasoft.UpdateListInput) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Furnitures_List: NewFurnituresService(repos.FurniruresList),
	}
}
