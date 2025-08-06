package companycontroller

import (
	"tenjin/back/internal/company"

	"github.com/gin-gonic/gin"
)

type companyController struct {
	companyService company.CompanyServiceInterface
}

type CompanyControllerInterface interface {
	RetrieveCompanyInfo(c *gin.Context)
}

func NewCompanyController(companyService company.CompanyServiceInterface) CompanyControllerInterface {
	return &companyController{
		companyService: companyService,
	}
}
