package uploadfiletestcontroller

import (
	"tenjin/back/internal/utils/filestores"
)

type FileTestController struct {
	storageService *filestores.FileStoreService
}

func NewFileTestController(storageService *filestores.FileStoreService) *FileTestController {
	return &FileTestController{
		storageService: storageService,
	}
}
