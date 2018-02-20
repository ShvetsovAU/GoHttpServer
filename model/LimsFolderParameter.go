package model

import "gopkg.in/mgo.v2/bson"

type LimsFolderParameter struct {
	Id 			bson.ObjectId
	FolderId	bson.ObjectId
	Name 		string
	Value		string
}

type LimsFolderParameterCollection struct {
	Items	[]LimsFolderParameter
}