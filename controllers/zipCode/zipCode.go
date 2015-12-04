package zipCode

import (
	"fmt"
	"github.com/nicolas.garaza/models/zipCode"
	"github.com/nicolas.garaza/repository/zipCodeRepository"
)

/*
	Constructor
*/
func New() *ZipCodeController {
	fmt.Println("Created Controller")

	r := zipCodeRepository.New()

	return &ZipCodeController{repository: r}
}

type ZipCodeController struct {
	repository *zipCodeRepository.ZipCodeRepository
}


func (z ZipCodeController) Get() ([]zipCode.ZipCode, error) {
	return z.repository.GetList()
}

func (z ZipCodeController) GetByCode(code string) []zipCode.ZipCode {
	return []zipCode.ZipCode{}
}

func (z ZipCodeController) Post(zipcode zipCode.ZipCode) {
	z.repository.Insert(zipcode)
}
