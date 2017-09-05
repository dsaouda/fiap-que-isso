package models

import "gopkg.in/mgo.v2/bson"

type Slide struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
	Value string `json:"value"`
	Group string `json:"group"`
	Order int `json:"order"`
}
