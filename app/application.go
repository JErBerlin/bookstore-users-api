package app

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine = gin.Default()
)

func StartApplication() {
	mapURLs()
	router.Run(":8080")
}
