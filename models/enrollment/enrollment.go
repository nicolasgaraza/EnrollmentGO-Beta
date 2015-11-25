package enrollment

import (

)

type Enrollment struct {
	Id int
}

func (e Enrollment) GetName() string {
	return "pepe"
}