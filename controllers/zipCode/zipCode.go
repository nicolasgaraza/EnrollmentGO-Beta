package zipCode

import (
	"github.com/nicolas.garaza/models/zipCode"
	"github.com/nicolas.garaza/repository/zipCodeRepository"
	"fmt"
)
/*
	Constructor
*/
func New () *ZipCodeController{
	fmt.Println("Created Controller");
	
	r:= zipCodeRepository.New()	
	
	return &ZipCodeController{repository : r}
} 


type ZipCodeController struct {
	repository *zipCodeRepository.ZipCodeRepository
}

func (z ZipCodeController) Get() ([]zipCode.ZipCode, error){
	return z.repository.GetList() 
}

func (z ZipCodeController) GetByCode(code string) []zipCode.ZipCode{
	return []zipCode.ZipCode{} 
}