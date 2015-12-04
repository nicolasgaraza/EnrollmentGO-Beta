package zipCode

import (
	"gopkg.in/mgo.v2/bson"
	//"encoding/json"
)

type ZipCode struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Code      string `json:"code" bson:"code"`
	State     string `json:"state" bson:"state"`
	Cities    []string `json:"cities" bson:"cities"`
	Available bool `json:"available" bson:"available"`
}
