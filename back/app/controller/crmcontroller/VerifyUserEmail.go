package crmcontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

// VerifyUserEmail godoc
// @Summary Vérifie l'email d'un utilisateur
// @Description Valide le token de vérification envoyé par email et met à jour la propriété EmailVerified de l'utilisateur
// @Tags user
// @Accept json
// @Produce json
// @Param token query string true "Token de vérification envoyé par email"
// @Success 200 {object} ginresponse.JsonFormatterSwag "Email vérifié avec succès"
// @Failure 400 {object} ginresponse.JsonFormatterSwag "Paramètres invalides ou token invalide/expiré"
// @Failure 500 {object} ginresponse.JsonFormatterSwag "Erreur interne lors de la vérification"
// @Router /user/verify-email [get]
func (c CrmController) VerifyUserEmail(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		logger.Wf("Token manquant dans la requête")
		ginresponse.BadRequest(ctx, "Token manquant", ginresponse.ErrorModel{
			Message: "Le token est requis",
			Type:    "Validation",
			Detail:  "Paramètre 'token' manquant",
		})
		return
	}

	err := c.userService.VerifyEmail(ctx, token, c.emailVerificationService)
	if err != nil {
		logger.Wf("Erreur lors de la vérification du token: %v", err)
		ginresponse.BadRequest(ctx, "Échec de la vérification de l'email", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "VerifyEmail",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	logger.Sf("Email vérifié avec succès pour le token: %s", token)
	ginresponse.Success(ctx, "Email vérifié avec succès", map[string]string{})
}
