package model

import "gopkg.in/mgo.v2/bson"

type SelectionPoint struct {
	Id 		bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	Name	string			`bson:"name" json:"name"`
}

type SelectionPointCollection struct {
	Items []SelectionPoint
}