package uploadfiletestcontroller

import (
	"tenjin/back/internal/filestores"
)

type FileTestController struct {
	storageService *filestores.FileStoreService
}

func NewFileTestController(storageService *filestores.FileStoreService) *FileTestController {
	return &FileTestController{
		storageService: storageService,
	}
}
