package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type ID = bson.ObjectId

func NewID() ID {
	return bson.NewObjectId()
}
