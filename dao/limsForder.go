package dao

import (
	"github.com/shvetsovau/GoHttpServer/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LimsFolderDao struct {
	limsFolders	*mgo.Collection
}

// implement Crud interface
func (dao *LimsFolderDao) GetById(entity interface{}, id interface{}) error {
	err := dao.limsFolders.FindId(model.ToId(id)).One(entity)
	return err
}

//Доделать! нет никакого gid!
func (dao *LimsFolderDao) GetAll(gid string) (model.LimsFolderCollection, error) {
	limsFolderCollection := model.LimsFolderCollection{}
	query := bson.M{"gid": bson.ObjectIdHex(gid)}
	err := dao.limsFolders.Find(query).All(&limsFolderCollection.Items)
	if err != nil {
		return limsFolderCollection, err
	}
	return limsFolderCollection, nil
}

// implement LimsFolderDao interface
func (dao *LimsFolderDao) InsertLimsFolder(limsFolder *model.LimsFolder) (interface{}, error) {
	if limsFolder.Id == "" {
		limsFolder.Id = bson.NewObjectId()
	}
	return limsFolder.Id, dao.limsFolders.Insert(limsFolder)
}