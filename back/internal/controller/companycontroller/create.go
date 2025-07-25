package companycontroller

import (
	"fmt"
	"tenjin/back/internal/insee"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

// Create godoc
// @Summary Crée une nouvelle company
// @Description Permet de créer une nouvelle company
// @Tags company
// @Accept json
// @Produce json
// @Param company body insee.CompanyCreateDto true "Informations de la company à créer"
// @Success 201 {object} ginresponse.JsonFormatterSwag "Company créée avec succès"
// @Failure 400 {object} ginresponse.JsonFormatterSwag "Erreur de validation"
// @Failure 500 {object} ginresponse.JsonFormatterSwag "Erreur interne"
// @Router /company/create [post]
func (p *companyController) Create(c *gin.Context) {
	var companyCreateDto insee.CompanyCreateDto
	if err := c.ShouldBindJSON(&companyCreateDto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	_, err := p.companyService.Create(c, &companyCreateDto)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.InternalServerError(c, err.Error(), err.Error())
		return
	}

	ginresponse.Created(c, "Company créé avec succès", []string{})
}
