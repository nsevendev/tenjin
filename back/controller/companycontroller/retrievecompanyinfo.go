package companycontroller

import (
	"fmt"
	"tenjin/back/internal/company"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

// Retrieve godoc
// @Summary Récupère les infos d'une entreprise via SIRET/SIREN
// @Description Utilise l'API INSEE pour retrouver les informations d'une entreprise
// @Tags company
// @Accept json
// @Produce json
// @Param data body company.CompanyRetrieveDto true "SIRET et SIREN de l'entreprise"
// @Success 200 {object} ginresponse.JsonFormatterSwag "Entreprise trouvée"
// @Failure 400 {object} ginresponse.JsonFormatterSwag "Paramètres invalides"
// @Failure 500 {object} ginresponse.JsonFormatterSwag "Erreur interne lors de l'appel à l'API INSEE"
// @Router /company/retrieve-infos [post]
func (cc *companyController) RetrieveCompanyInfo(c *gin.Context) {
	var dto company.CompanyRetrieveDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	companyInfo, err := cc.companyService.RetrieveCompanyInfo(c.Request.Context(), dto.Siret, dto.Siren)
	if err != nil {
		logger.Ef("Erreur lors de la recuperation de l'entreprise : %v", err)
		ginresponse.InternalServerError(c, "Erreur lors de la recuperation des informations INSEE", err.Error())
		return
	}

	ginresponse.Success(c, "Entreprise trouvee", companyInfo)
}
