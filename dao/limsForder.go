package dao

import (
	"github.com/shvetsovau/gohttpserver/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LimsFolderDao struct {
	limsFolders	*mgo.Collection
}

func (dao *LimsFolderDao) GetById(entity interface{}, id interface{}) error {
	err := dao.limsFolders.FindId(model.ToId(id)).One(entity)
	return err
}

// implement Crud interface
func (dao *LimsFolderDao) GetLimsFolder(limsFolderId interface{}) (*model.LimsFolder, error) {
	limsFolder := &model.LimsFolder{}
	err := dao.GetById(limsFolder, limsFolderId)
	return limsFolder, err
}

func (dao *LimsFolderDao) GetAll() (model.LimsFolderCollection, error) {
	limsFolderCollection := model.LimsFolderCollection{}
	query := bson.M{}
	err := dao.limsFolders.Find(query).All(&limsFolderCollection.Items)
	if err != nil {
		return limsFolderCollection, err
	}
	return limsFolderCollection, nil
}

// implement LimsFolderDao interface
func (dao *LimsFolderDao) Insert(limsFolder *model.LimsFolder) (interface{}, error) {
	if limsFolder.Id == "" {
		limsFolder.Id = bson.NewObjectId()
	}
	return limsFolder.Id, dao.limsFolders.Insert(limsFolder)
}

// implement LimsFolderDao interface
func (dao LimsFolderDao) Update(limsFolder *model.LimsFolder) error {
	_, err := dao.limsFolders.UpsertId(limsFolder.Id, limsFolder)
	return err
}

func (dao LimsFolderDao) Delete(id interface{}) error {
	return dao.limsFolders.Remove(bson.M{"_id": model.ToId(id)})
}