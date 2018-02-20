package model

import "gopkg.in/mgo.v2/bson"

type MaterialType struct {
	Id			bson.ObjectId
	Name		string
	Materials	MaterialCollection
}