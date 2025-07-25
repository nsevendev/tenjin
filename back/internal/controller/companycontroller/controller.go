package companycontroller

import (
	"tenjin/back/internal/insee"
)

type companyController struct {
	companyService insee.CompanyServiceInterface
}

type CompanyControllerInterface interface {
}

func NewCompanyController(companyService insee.CompanyServiceInterface) CompanyControllerInterface {
	return &companyController{
		companyService: companyService,
	}
}
