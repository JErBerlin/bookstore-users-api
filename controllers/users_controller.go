package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "Hit CreateUser")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "Hit GetUser")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusOK, "Hit SearchUser")
}
