package companycontroller

import (
	"tenjin/back/internal/company"
)

type companyController struct {
	companyService company.CompanyServiceInterface
}

type CompanyControllerInterface interface {

}

func NewCompanyController(companyService company.CompanyServiceInterface) CompanyControllerInterface {
	return &companyController{
		companyService: companyService,
	}
}
