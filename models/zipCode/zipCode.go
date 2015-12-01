package zipCode

import (
 	"gopkg.in/mgo.v2/bson"
)

type ZipCode struct {
	Id bson.ObjectId
	Code string
	State string
	Cities []string
	Available bool
}

