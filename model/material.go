package model

import "gopkg.in/mgo.v2/bson"

type Material struct {
	Id				bson.ObjectId
	Name			string
	MaterialTypeId	bson.ObjectId
}

type MaterialCollection struct {
	Items []Material
}