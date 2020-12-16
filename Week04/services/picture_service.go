package services

import (
	//"Week04/serializer"
	"Week04/dao"
	"Week04/models"
)

type PictureService struct {
	dao *dao.PictureDao
}

func NewPictureService() *PictureService {
	return &PictureService{
		dao: dao.NewPictureDao(dao.InstanceDB()),
	}
}

func (s *PictureService) Query(id int) (string, error) {
	return s.dao.Query(id)
}


func (s *PictureService) GetAllPicture() ([]*models.Picture, error) {
	return s.dao.GetAllPicture()
}
