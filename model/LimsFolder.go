package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type LimsFolder struct {
	Id 			bson.ObjectId 	`bson:"_id,omitempty" json:"-"`
	FolderId 	string        	`bson:"order_id" json:"folderId"`
	FolderName 	string        	`bson:"order_name" json:"folderName"`
	FolderDate	time.Time 		`bson:"date" json:"folderDate"`
}