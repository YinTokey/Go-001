package services

import (
	//"Week04/serializer"
	"Week04/dao"
	"Week04/models"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewPictureService, dao.Provider)

type PictureService struct {
	dao *dao.PictureDao
}

func NewPictureService(dao *dao.PictureDao) *PictureService {
	return &PictureService{
		dao: dao,
	}
}

//func NewPictureService() *PictureService {
//	return &PictureService{
//		dao: dao.NewPictureDao(dao.InstanceDB()),
//	}
//}

func (s *PictureService) Query(id int) (string, error) {
	return s.dao.Query(id)
}


func (s *PictureService) GetAllPicture() ([]*models.Picture, error) {
	return s.dao.GetAllPicture()
}
