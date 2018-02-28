package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type LimsFolder struct {
	Id 					bson.ObjectId 					`bson:"_id,omitempty" json:"id"`
	FolderId 			string        					`bson:"folderId" json:"folderId"`
	FolderName 			string        					`bson:"folderName" json:"folderName"`
	FolderDate			time.Time 						`bson:"folderDate" json:"folderDate"`
	MaterialTypeId		bson.ObjectId					`bson:"materialTypeId" json:"materialTypeId"`
	MaterialId			bson.ObjectId					`bson:"materialId" json:"materialId"`

	//provider or customer
	ContractorId		bson.ObjectId					`bson:"contractorId" json:"contractorId"`
	SelectionPointId	bson.ObjectId					`bson:"selectionPointId" json:"selectionPointId"`
	Parameters 			LimsFolderParameterCollection	`bson:"parameters" json:"parameters"`
}

type LimsFolderCollection struct {
	Items []LimsFolder `json:"items"`
}