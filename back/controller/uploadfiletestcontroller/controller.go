package uploadfiletestcontroller

import (
	"tenjin/back/internal/storage"

	"github.com/gin-gonic/gin"
)

type fileTestController struct {
	storageService *storage.Service
}

type FileTestControllerInterface interface {
	Create(c *gin.Context)
}

func NewFileTestController(storageService *storage.Service) FileTestControllerInterface {
	return &fileTestController{
		storageService: storageService,
	}
}
