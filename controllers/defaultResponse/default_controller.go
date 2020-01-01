package defaultResponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultResponse(c *gin.Context) {
	c.String(http.StatusOK, "Happy new year!")
}
