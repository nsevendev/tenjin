package companycontroller

import (
	"tenjin/back/internal/insee"

	"github.com/gin-gonic/gin"
)

type companyController struct {
	companyService insee.CompanyServiceInterface
}

type CompanyControllerInterface interface {
	Create(*gin.Context)
}

func NewCompanyController(companyService insee.CompanyServiceInterface) CompanyControllerInterface {
	return &companyController{
		companyService: companyService,
	}
}
