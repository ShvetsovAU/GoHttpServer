package model

import "gopkg.in/mgo.v2/bson"

type Contractor struct {
	Id		bson.ObjectId
	Name	string
	Points	SelectionPointCollection
}