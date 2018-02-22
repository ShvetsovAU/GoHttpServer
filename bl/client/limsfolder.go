package client

import (
	"github.com/shvetsovau/gohttpserver/dao"
	"github.com/shvetsovau/gohttpserver/model"
)

type LimsFolderBL struct {
	BaseBL
	dao *dao.LimsFolderDao
}

func (bl *LimsFolderBL) Create(bill *model.LimsFolder, gid string) *model.LimsFolder {
	//utils.PanicIfErr(bl.permit.Check(rights.GroupId(gid)), billBlTAG)

	//bill.GroupID = model.ToId(gid)
	_, err := bl.dao.InsertBill(bill)
	//utils.PanicIfErr(err, billBlTAG)

	// обновить данные Schedule для данного счета
	//bl.updateScheduleFor(bill)

	//bl.updateStatus(bill)
	return bill
}