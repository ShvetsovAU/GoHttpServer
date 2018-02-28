package model

import "gopkg.in/mgo.v2/bson"

type Contractor struct {
	Id		bson.ObjectId				`bson:"_id,omitempty" json:"id"`
	Name	string						`bson:"name" json:"name"`
	Points	SelectionPointCollection	`bson:"points" json:"points"`
}