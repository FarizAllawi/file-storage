package main

import (
	// "fmt"
	// "eca_file_storage/config"
	"github.com/gin-contrib/cors"
	"eca_file_storage/routes"
)

func main() {

	router := routes.SetupRoutes()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS", "TRACE", "CONNECT"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Date" , "Content-Length"},
	}))
	
	router.Run(":8080")
	// router.RunTLS(":8080", "./certificates/STAR_sakafarma_com.pem", "./certificates/commercial.key")
}
