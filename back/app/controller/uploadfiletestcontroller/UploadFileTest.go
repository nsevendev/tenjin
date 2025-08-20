package uploadfiletestcontroller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"tenjin/back/internal/utils/filestores"
)

// UploadFileTest godoc
// @Summary Enregistrement d'un fichier dans le stockage R2
// @Description Test enregistre un fichier dans le stockage R2
// @Tags UploadFileTest
// @Accept json
// @Produce json
// @Success 200 {object} ginresponse.JsonFormatterSwag "Fichier enregistré avec succès"
// @Failure 400 {object} ginresponse.JsonFormatterSwag "Erreur à la récupération du fichier"
// @Failure 500 {object} ginresponse.JsonFormatterSwag "Erreur à l'enregistrement dans S3"
// @Router /uploadfile/test [post]
func (cc *FileTestController) UploadFileTest(c *gin.Context) {
	fileData, err := cc.storageService.ExtractFileFromForm(c, "file")
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.BadRequest(c, "Erreur à la récupération du fichier", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "error",
			Detail:  fmt.Sprintf("%v", err),
		})
	}

	// envoi sur R2 avec le nom original comme clé
	result, err := cc.storageService.UploadBytes(context.Background(), "test", fileData.Filename, fileData.Data)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.BadRequest(c, "Erreur à l'enregistrement dans S3", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "error",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	type Test struct {
		Result filestores.UploadStoreResult
	}

	res := Test{
		Result: *result,
	}

	ginresponse.Success(c, "Fichier enregistré avec succès", res)
}
