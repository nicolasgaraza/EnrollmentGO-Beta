package repository

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

func New() *session {

	fmt.Println("New Session")
	
	_s, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("There was an error");
		panic(err)
	}

	_s.SetMode(mgo.Monotonic, true)
	
	s:= &session{session : _s}
	
	return s
}


type session struct {
	session *mgo.Session
} 

func ( s session )DataBase() *mgo.Database{
	return s.session.DB("test")
} 

func ( s session) Close() {
	s.session.Close()
}

type SessionInterface interface{
	DataBase() *mgo.Database
	Close()
}
