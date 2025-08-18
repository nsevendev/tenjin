package uploadfiletestcontroller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"io"
	"net/http"
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
func (cc *fileTestController) Create(c *gin.Context) {
	// Récupération du fichier depuis le form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fichier manquant"})
		return
	}

	f, err := file.Open()
	if err != nil {
		logger.Ef("impossible d'ouvrir le fichier: %v", err)
		ginresponse.BadRequest(c, "impossible d'ouvrir le fichier", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "error",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		logger.Ef("lecture du fichier échouée: %v", err)
		ginresponse.BadRequest(c, "lecture du fichier échouée", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "error",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	// envoi sur R2 avec le nom original comme clé
	if _, err := cc.storageService.UploadBytes(context.Background(), "test", file.Filename, data); err != nil {
		logger.Ef("enregistrement sur r2 échouée: %v", err)
		ginresponse.BadRequest(c, "enregistrement sur r2 échouée", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "error",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	type Test struct {
		Filename string `json:"filename"`
		Size     int64  `json:"size"`
	}

	res := Test{
		Filename: file.Filename,
		Size:     file.Size,
	}

	ginresponse.Success(c, "Entreprise créée avec succès", res)
}
