package app

import (
	"github.com/JErBerlin/bookstore-users-api/controllers"
)

func mapURLs() {
	router.GET("/ping", controllers.Ping)
	router.GET("/", controllers.DefaultResponse)
}
