package routes

import (
	"github.com/gin-gonic/gin"
	media "eca_file_storage/controller/media" // import media controller

)

func SetupRoutes() *gin.Engine {
	route := gin.Default()
	
	// Routess heree
	files := route.Group("/files")
	{
		files.POST("upload", media.UploadFile)
		files.GET("get", media.GetFile)
		files.DELETE("delete", media.DeleteFile)
	}


	return route
	// Booting up the server
	// router.Run("localhost:8080")
}
