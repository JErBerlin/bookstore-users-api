package app

import (
	"github.com/JErBerlin/bookstore_users-api/controllers/defaultResponse"
	"github.com/JErBerlin/bookstore_users-api/controllers/ping"
	"github.com/JErBerlin/bookstore_users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)
	router.GET("/", defaultResponse.DefaultResponse)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
