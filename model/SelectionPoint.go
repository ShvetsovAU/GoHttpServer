package model

import "gopkg.in/mgo.v2/bson"

type SelectionPoint struct {
	Id 		bson.ObjectId
	Name	string
}

type SelectionPointCollection struct {
	Items []SelectionPoint
}