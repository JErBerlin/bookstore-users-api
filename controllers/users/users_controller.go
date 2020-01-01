package users

import (
	"fmt"
	"github.com/JErBerlin/bookstore-users-api/domain/users"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	log.Print(err)
	fmt.Println(string(bytes))
	if string(bytes) == "" {
		bytes = []byte("No info of the user available")
	}
	c.String(http.StatusOK, string(bytes))
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Hit GetUser")
}

/*
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Hit SearchUser")
}
*/
