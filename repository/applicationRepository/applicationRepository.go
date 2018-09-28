package applicationRepository

import (
	"github.com/nicolas.garaza/repository/redisSession"
	"gopkg.in/redis.v3"
	"encoding/json"
	"github.com/nicolas.garaza/models/enrollment"
)

type ApplicationRepository struct {
	session *redis.Client 
}

func New() *ApplicationRepository {
	r := redisSession.New()
	return &ApplicationRepository{session: r}
}



func (a ApplicationRepository) getApplication(key string) (string, error){
	 
	 val , er := a.session.Get(key).Result()
	 
	 return val, er
	 
}

func (a ApplicationRepository) saveApplication(app enrollment.Application) (error){
	 
	 b , _ := json.Marshal(app)
	 
	 er := a.session.Set(app.Id, b /*convert to string*/, 0).Err() 
	 
	 return  er
	 
}


func (a ApplicationRepository) Get (code string){
	v , er := a.getApplication(code)
	
	if er != nil{
		panic(er)
	}
	var app enrollment.Application
	b := []byte(v)
	json.Unmarshal(b, &app)
}

func (a ApplicationRepository) Save(enrollment.Application){
	
}