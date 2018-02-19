package dao

import (
	"github.com/shvetsovau/GoHttpServer/model"
	"gopkg.in/mgo.v2"
)

type LimsFolderDao struct {
	limsFolders	*mgo.Collection
}

// implement Crud interface
func (dao *LimsFolderDao) GetById(entity interface{}, id interface{}) error {
	err := dao.limsFolders.FindId(model.ToId(id)).One(entity)
	return err
}