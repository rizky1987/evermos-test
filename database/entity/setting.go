package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type Setting struct {
	Id               		*bson.ObjectId 			`bson:"_id,omitempty"`
	Key            			string         			`bson:"key"`
	Value         		    string         			`bson:"value"`
}