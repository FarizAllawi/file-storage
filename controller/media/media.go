// controller/media/media.go

package media

import (
	"os"
	"fmt"	
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func UploadFile(context *gin.Context ) {
	var fileName, fileNameUUID, contentType string
	var fileSize int64

	// 1. parse input, type multipart/form-data
	form , _ := context.MultipartForm()	

	// 2. retrieve file from posted form-data
	files := form.File["Files"]
	if files == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Files is required"})
		return 
	}
	
	for _, file := range files {
		fileNameSplit := strings.Split(file.Filename, ".")

		id := uuid.New()
		fileName = file.Filename
		fileNameUUID = fmt.Sprintf("%s."+fileNameSplit[len(fileNameSplit)-1] , id) 
		contentType = file.Header["Content-Type"][0]
		fileSize = file.Size

		// Save File to storage
		err := context.SaveUploadedFile(file, "storage/"+fileName)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		// Rename uploaded File
		os.Rename("storage/"+file.Filename , "storage/"+fileNameUUID)
	}

	context.JSON(http.StatusOK, gin.H{
		"success": "Upload successfully",
		"data": gin.H{
			"fileName": fileName,
			"filePath": "storage/"+fileNameUUID,
			"contentType": contentType,
			"fileSize": fileSize,
		},
	})
}

func GetFile(context *gin.Context) {	
	var filePath string
	queryParams := context.Request.URL.Query()

	if len(queryParams) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "filePath is required"})
		return 
	}

	filePath = queryParams["filePath"][0]

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}

	contentType := http.DetectContentType(file[:512])

	// context.Header("Content-Disposition", "attachment; filename="+strings.Split(filePath ,"/")[1])
	context.Data(http.StatusOK, contentType, file)
}

func DeleteFile(context *gin.Context) {
	var filePath string
	queryParams := context.Request.URL.Query()

	if len(queryParams) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "filePath is required"})
		return 
	}

	filePath = queryParams["filePath"][0]

	// Removing file
    // Using Remove() function
	err := os.Remove(filePath)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	context.JSON(http.StatusOK, gin.H{
		"success": "Delete successfully",
	})
}