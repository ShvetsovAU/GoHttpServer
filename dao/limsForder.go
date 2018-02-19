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

func (dao *LimsFolderDao) GetAll(gid string) (model.LimsFoldersCollection, error) {
	limsFoldersCollection := model.LimsFoldersCollection{}
	query := bson.M{"gid": bson.ObjectIdHex(gid)}
	err := dao.limsFolders.Find(query).All(&limsFoldersCollection.Items)
	if err != nil {
		return limsFoldersCollection, err
	}
	return limsFoldersCollection, nil
}

// implement LimsFolderDao interface
func (dao *LimsFolderDao) InsertBill(limsFolder *model.LimsFolder) (interface{}, error) {
	if limsFolder.Id == "" {
		limsFolder.Id = bson.NewObjectId()
	}
	return limsFolder.Id, dao.limsFolders.Insert(limsFolder)
}