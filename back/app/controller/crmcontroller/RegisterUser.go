package crmcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"tenjin/back/internal/crm"
)

func (c CrmController) RegisterUser(ctx *gin.Context) {
	var userCreateDto crm.UserCreateDto
	if err := ctx.ShouldBindJSON(&userCreateDto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(ctx, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	user, err := c.userService.CreateUser(ctx, userCreateDto)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.BadRequest(ctx, "Erreur à la création de l'utilisateur", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "CreateUser",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	logger.Sf("Utilisateur créé avec succès: %s", user.Email)
	ginresponse.Success(ctx, "Utilisateur créé avec succès", map[string]string{})
}
