package authcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"tenjin/back/internal/auth"
)

func (a *AuthController) Login(ctx *gin.Context) {
	var loginDto auth.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		logger.Ef("%v", err)
		ginresponse.BadRequest(ctx, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	user, err := a.userService.FindByEmail(ctx, loginDto.Email)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.InternalServerError(ctx, "Erreur identifiant", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "MongoDB",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	token, err := a.authService.CreateToken(user, loginDto)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.InternalServerError(ctx, "Erreur creation token", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Authentication",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	// config cookie
	cookieName := a.authService.NameCookie()
	maxAge := 24 * 60 * 60 // 24 heures en secondes

	// déterminer si on est en HTTPS ou HTTP
	secure := ctx.Request.Header.Get("X-Forwarded-Proto") == "https" ||
		ctx.Request.TLS != nil ||
		gin.Mode() == gin.ReleaseMode

	// créer le cookie HttpOnly
	ctx.SetCookie(
		cookieName, // name
		token,      // value
		maxAge,     // maxAge (en secondes)
		"/",        // path
		"",         // domain (vide = domaine actuel)
		secure,     // secure (HTTPS uniquement en production)
		true,       // httpOnly (très important pour la sécurité)
	)

	// ajout des headers de sécurité
	ctx.Header("Set-Cookie", fmt.Sprintf("%s=%s; Path=/; Max-Age=%d; HttpOnly; SameSite=Strict%s",
		cookieName, token, maxAge, func() string {
			if secure {
				return "; Secure"
			}
			return ""
		}()))

	ginresponse.Success(ctx, "Connexion réussie", map[string]string{})
}
