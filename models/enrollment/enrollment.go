package enrollment

import (
	
)

type Enrollment struct {
	Id int
	Applicant PersonData
	Dependant PersonData
	Address Address
}

func (e Enrollment) GetName() string {
	return "pepe"
}