package app

import (
	"github.com/JErBerlin/bookstore-users-api/controllers"
)

func mapURLs() {
	router.GET("/ping", controllers.Ping)
	router.GET("/", controllers.DefaultResponse)

	router.GET("/users/search", controllers.SearchUser)
	// router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
}
