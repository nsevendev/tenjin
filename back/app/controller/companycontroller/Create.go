package companycontroller

import (
	"fmt"
	"tenjin/back/internal/company"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

// Create godoc
// @Summary Crée une entreprise
// @Description Crée une entreprise dans la base de données MongoDB
// @Tags company
// @Accept json
// @Produce json
// @Param data body company.CompanyCreateDto true "Informations de l'entreprise à créer"
// @Success 201 {object} ginresponse.JsonFormatterSwag "Entreprise créée"
// @Failure 400 {object} ginresponse.JsonFormatterSwag "Paramètres invalides"
// @Failure 500 {object} ginresponse.JsonFormatterSwag "Erreur interne lors de la création"
// @Router /company/register [post]
func (cc *companyController) Create(c *gin.Context) {
	var dto company.CompanyCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	newCompany, err := cc.companyService.Create(c.Request.Context(), dto)
	if err != nil {
		logger.Ef("Erreur lors de la création de l'entreprise : %v", err)
		ginresponse.InternalServerError(c, "Erreur lors de la création de l'entreprise", err.Error())
		return
	}

	ginresponse.Created(c, "Entreprise créée avec succès", newCompany)
}
