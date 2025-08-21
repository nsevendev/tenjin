package router

import (
	"github.com/gin-gonic/gin"
	"tenjin/back/app/controller/uploadfiletestcontroller"
	"tenjin/back/internal/utils/filestores"
)

func RegisterUploadFileTest(v1 *gin.RouterGroup, deps *dependencies) {
	fileStoreService := filestores.NewService(deps.R2Adapter, filestores.FileStoreConfig{
		KeyPrefix:      "tenjin/uploads/",
		MaxSize:        10 * 1024 * 1024, // 10 Mo
		AllowedMIMEs:   []string{"image/jpeg", "image/png", "application/pdf"},
		UseDateFolders: true, // passé à false pour une autre config si pas de dossier date
	})

	fileTestController := uploadfiletestcontroller.NewFileTestController(fileStoreService)

	fileTestGroup := v1.Group("/uploadfile")
	{
		fileTestGroup.POST("/test", fileTestController.UploadFileTest)
	}
}
