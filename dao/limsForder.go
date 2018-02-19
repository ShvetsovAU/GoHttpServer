package dao

import (
	"github.com/shvetsovau/GoHttpServer/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LimsFolderDao struct {
	limsFolders	*mgo.Collection
	utilities 	*mgo.Collection
}

// implement Crud interface
func (dao *LimsFolderDao) GetById(entity interface{}, id interface{}) error {
	err := dao.limsFolders.FindId(model.ToId(id)).One(entity)
	return err
}