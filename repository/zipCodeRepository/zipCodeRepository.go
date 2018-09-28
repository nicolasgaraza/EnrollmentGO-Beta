package zipCodeRepository

import (
	//"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	"github.com/nicolas.garaza/repository"
	"github.com/nicolas.garaza/models/zipCode"
	"fmt"
)

type ZipCodeRepository struct {
	session repository.SessionInterface 
}

func New () *ZipCodeRepository{
	return &ZipCodeRepository {  session : repository.New() } 
}

func (z ZipCodeRepository) GetList() ([]zipCode.ZipCode, error) {
	var zipCodes []zipCode.ZipCode
	
	defer z.session.Close()
	
	er:= z.session.DataBase().C("zipcodes").Find(nil).All(&zipCodes)
	
	fmt.Printf("%+v\n", zipCodes)
	
	if er == nil{
		return zipCodes, nil
	}
	return nil, er
}

func (z ZipCodeRepository) GetByCode (code string) (zipCode.ZipCode, error) {
	
	var zipcode zipCode.ZipCode
	
	defer z.session.Close()
	
	er := z.session.DataBase().C("zipcodes").Find( zipCode.ZipCode { Code : code }).One(&zipcode)
	
	if er != nil{
		return zipcode, nil
	}
	return zipCode.ZipCode{}, er
}

func (z ZipCodeRepository) Insert ( zipcode zipCode.ZipCode){
	
	zipcode.Id = bson.NewObjectId()
	
	defer z.session.Close()
	
	er := z.session.DataBase().C("zipcodes").Insert(&zipcode)
	
	if er != nil{
		panic(er)
	}
	return
}



