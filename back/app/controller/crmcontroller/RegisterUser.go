package crmcontroller

import (
	"fmt"
	"net/http"
	"tenjin/back/internal/crm"
	"tenjin/back/internal/jobs"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
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

	token, err := c.emailVerificationService.GenerateToken(user.ID)
	if err != nil {
		logger.Ef("Erreur lors de la génération du token pour %s : %v", user.Email, err)
		ginresponse.Error(ctx, http.StatusInternalServerError, "Impossible de générer un token de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "GenerateToken",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	verificationURL := fmt.Sprintf("https://%s/user/verify-email?token=%s", ctx.Request.Host, token)

	job := jobs.Job{
		Name:     "mail:send",
		MaxRetry: 3,
		Payload: map[string]string{
			"user_id": user.ID.Hex(),
			"email":   user.Email,
			"subject": "Vérifiez votre adresse email",
			"body":    fmt.Sprintf(
				"Bonjour %s,\n\nVotre inscription à Tenjin a bien été enregistrée.\n\nVeuillez vérifier votre adresse email en cliquant sur le lien suivant :\n%s",
				user.Firstname,
				verificationURL,
			),
		},
	}
	jobs.ProcessJob(ctx, job)

	logger.Sf("Utilisateur créé avec succès: %s", user.Email)
	ginresponse.Success(ctx, "Utilisateur créé avec succès", map[string]string{})
}
