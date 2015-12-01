package repository

import (
	"gopkg.in/mgo.v2"
   // "gopkg.in/mgo.v2/bson"
   "fmt"
)

func NewSession() *mgo.Database {
	
	fmt.Println("New Session")
	
	 session, err := mgo.Dial("localhost")
	  if err != nil {
                panic(err)
        }
        defer session.Close()
        
        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
        
        return session.DB("test")
} 