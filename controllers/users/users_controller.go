package users

import (
	"encoding/json"
	"fmt"
	"github.com/JErBerlin/bookstore-users-api/domain/users"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
		c.String(http.StatusBadRequest, string(bytes))
		return
	}
	if err = json.Unmarshal(bytes, &user); err != nil {
		log.Print(err)
		c.String(http.StatusUnprocessableEntity, string(bytes))
		return
	}
	if user == users.NewUser(0, "", "", "", "") {
		c.String(http.StatusUnprocessableEntity, string(bytes))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%#v", user))
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Hit GetUser")
}
