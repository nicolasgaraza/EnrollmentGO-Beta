package enrollment

import (
	"github.com/nu7hatch/gouuid"
)

type Application struct {
	Id string
}

func (a Application) NewApplication() Application {
	
	u, _ := uuid.NewV4()
	
	return Application {Id : u.String()}
}