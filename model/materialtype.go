package model

import "gopkg.in/mgo.v2/bson"

type MaterialType struct {
	Id			bson.ObjectId			`bson:"_id,omitempty" json:"id"`
	Name		string					`bson:"name" json:"name"`
	Materials	MaterialCollection		`bson:"materialItems" json:"materialItems"`
}