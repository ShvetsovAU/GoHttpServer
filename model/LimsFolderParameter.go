package model

import "gopkg.in/mgo.v2/bson"

type LimsFolderParameter struct {
	Id 			bson.ObjectId		`bson:"_id,omitempty" json:"id"`
	FolderId	bson.ObjectId		`bson:"folderId,omitempty" json:"folderId"`
	Name 		string				`bson:"name" json:"name"`
	Value		string				`bson:"value" json:"value"`
}

type LimsFolderParameterCollection struct {
	Items	[]LimsFolderParameter
}