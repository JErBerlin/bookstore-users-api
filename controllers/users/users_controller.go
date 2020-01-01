package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Hit CreateUser")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Hit GetUser")
}

/*
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Hit SearchUser")
}
*/
