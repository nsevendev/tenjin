package usercontroller

import (
	"fmt"
	"tenjin/back/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

// Create godoc
// @Summary Crée un utilisateur
// @Description Crée un utilisateur dans la base de données MongoDB
// @Tags user
// @Accept json
// @Produce json
// @Param data body user.UserCreateDto true "Informations de l'utilisateur à créer"
// @Success 201 {object} ginresponse.JsonFormatterSwag "Utilisateur créé"
// @Failure 400 {object} ginresponse.JsonFormatterSwag "Paramètres invalides"
// @Failure 500 {object} ginresponse.JsonFormatterSwag "Erreur interne lors de la création"
// @Router /user/register [post]
func (uc *userController) Create(c *gin.Context) {
	var dto user.UserCreateDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	newUser, err := uc.userService.Create(c.Request.Context(), dto)
	if err != nil {
		logger.Ef("Erreur lors de la creation de l'utilisateur : %v", err)
		ginresponse.InternalServerError(c, "Erreur lors de la creation de l'utilisateur", err.Error())
		return
	}

	ginresponse.Created(c, "Utilisateur cree avec succes", newUser)
}
