package enrollment

import (
	"gopkg.in/mgo.v2/bson"
)

type Enrollment struct {
	Id bson.ObjectId
	EnrollmentNumber int
	Applicant PersonData
	Dependant PersonData
	Address Address
}

func (e Enrollment) GetName() string {
	return "pepe"
}