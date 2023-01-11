package config 

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)


// Config for CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Max-Age", "0") //"86400"
		context.Writer.Header().Set("Access-Control-Allow-Methods", "*" ) // "POST, GET, OPTIONS, PUT, DELETE, UPDATE"
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*" ) //"Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
		context.Writer.Header().Set("Access-Control-Expose-Headers", "*") // Content-Length, multipart/form-data, Content-Disposition
		// context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}