package service

import (
	restapimediasoft "tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft"
	"tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft/internal/repository"
)

type FurnituresListService struct {
	repo repository.FurniruresList
}

func NewFurnituresService(repo repository.FurniruresList) *FurnituresListService {
	return &FurnituresListService{repo: repo}
}
func (s *FurnituresListService) Create(list restapimediasoft.Furnitures_list) (int, error) {
	return s.repo.Create(list)
}
func (s *FurnituresListService) GetAll() ([]restapimediasoft.Furnitures_list, error) {
	return s.repo.GetAll()
}
func (s *FurnituresListService) GetByID(listID int) (restapimediasoft.Furnitures_list, error) {
	return s.repo.GetByID(listID)
}
func (s *FurnituresListService) DeleteList(listID int) error {
	return s.repo.DeleteList(listID)
}
func (s *FurnituresListService) UpdateAll(listID int, input restapimediasoft.UpdateListInput) error {
	return s.repo.UpdateAll(listID, input)
}
func (s *FurnituresListService) Update(listID int, input restapimediasoft.UpdateListInput) error {
	return s.repo.Update(listID, input)
}
