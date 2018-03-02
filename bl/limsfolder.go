package bl

import (
	"github.com/shvetsovau/gohttpserver/dao"
	"github.com/shvetsovau/gohttpserver/model"
	"github.com/shvetsovau/gohttpserver/common/utils"
)

type LimsFolderBL struct {
	BaseBL
	dao *dao.LimsFolderDao
}

const (
	limsFolderBlTAG	= "LimsFolderBL"
)

func InitDao(bl *LimsFolderBL) {

	if bl.dao == nil {
		bl.dao = dao.GetLimsFolderDao()
	}
}

func (bl *LimsFolderBL) GetAll() *model.LimsFolderCollection {

	InitDao(bl)

	limsFolderCollection, err := bl.dao.GetAll()
	utils.PanicIfErr(err)

	return &limsFolderCollection
}

func (bl *LimsFolderBL) GetLimsFolder(id string) *model.LimsFolder {

	InitDao(bl)

	limsFolder, err := bl.dao.GetLimsFolder(id)
	utils.PanicIfErr(err)

	return limsFolder
}

func (bl *LimsFolderBL) Create(limsFolder *model.LimsFolder) *model.LimsFolder {

	InitDao(bl)

	_, err := bl.dao.Insert(limsFolder)
	utils.PanicIfErr(err)

	return limsFolder
}

func (bl *LimsFolderBL) Update(limsFolder *model.LimsFolder) {

	InitDao(bl)

	if limsFolder.Id == "" {
		// insert
		_, err := bl.dao.Insert(limsFolder,)
		utils.PanicIfErr(err,limsFolderBlTAG)

	} else {
		//update
		utils.PanicIfErr(bl.dao.Update(limsFolder), limsFolderBlTAG)
		}
}

func (bl *LimsFolderBL) Delete(id string) {

	InitDao(bl)

	utils.PanicIfErr(bl.dao.Delete(id), limsFolderBlTAG)
}