package router

import (
	"github.com/gin-gonic/gin"
	"tenjin/back/controller/uploadfiletestcontroller"
	"tenjin/back/internal/storage"
)

func RegisterUploadFileTest(v1 *gin.RouterGroup, deps *Dependencies) {
	storageService := storage.NewService(deps.R2Adapter, storage.Config{
		KeyPrefix:      "tenjin/uploads/",
		MaxSize:        10 * 1024 * 1024, // 10 Mo
		AllowedMIMEs:   []string{"image/jpeg", "image/png", "application/pdf"},
		UseDateFolders: true,
	})

	fileTestController := uploadfiletestcontroller.NewFileTestController(storageService)

	fileTestGroup := v1.Group("/uploadfile")
	{
		fileTestGroup.POST("/test", fileTestController.Create)
	}
}
