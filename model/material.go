package model

import "gopkg.in/mgo.v2/bson"

type Material struct {
	Id				bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	Name			string			`bson:"name" json:"name"`
	MaterialTypeId	bson.ObjectId	`bson:"materialId,omitempty" json:"materialId"`
}

type MaterialCollection struct {
	Items []Material
}