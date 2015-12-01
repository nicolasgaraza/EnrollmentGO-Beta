package enrollment

import (
	"time"
)

type PersonData struct {
	LastName string
	FirstName string
	SSN string
	DateOfBirth time.Time
}
