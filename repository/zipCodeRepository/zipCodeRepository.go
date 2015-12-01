package zipCodeRepository

import (
	"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"*/
	"github.com/nicolas.garaza/repository"
	"github.com/nicolas.garaza/models/zipCode"
)

type ZipCodeRepository struct {
	database *mgo.Database 
}

func New () *ZipCodeRepository{
	return &ZipCodeRepository {  database : repository.NewSession() } 
}

func (z ZipCodeRepository) GetList() ([]zipCode.ZipCode, error) {
	var zipCodes []zipCode.ZipCode
	
	er:= z.database.C("zipcodes").Find(nil).All(&zipCodes)
	
	if er != nil{
		return zipCodes, nil
	}
	return nil, er
}

func (z ZipCodeRepository) GetByCode (code string) (zipCode.ZipCode, error) {
	
	var zipcode zipCode.ZipCode
	
	er := z.database.C("zipcodes").Find( zipCode.ZipCode { Code : code }).One(&zipcode)
	
	if er != nil{
		return zipcode, nil
	}
	return zipCode.ZipCode{}, er
}



