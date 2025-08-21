package authcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
)

func (a *AuthController) Logout(ctx *gin.Context) {
	cookieName := a.authService.NameCookie()

	// Supprimer le cookie en définissant maxAge à -1
	ctx.SetCookie(
		cookieName,
		"", // valeur vide
		-1, // maxAge négatif = suppression
		"/",
		"",
		false, // secure (pas nécessaire pour la suppression)
		true,  // httpOnly
	)

	ginresponse.Success(ctx, "Déconnexion réussie", map[string]string{
		"message": "Vous avez été déconnecté avec succès.",
	})
}
