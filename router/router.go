package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize router
	router := gin.Default()
	//Initialize Routes
	initializeRoutes(router)
	//Define Port
	router.Run(":8080")
}
