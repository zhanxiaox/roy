package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	code int
	msg  string
	data string
}

type request struct {
	cmd   string
	token string
	data  string
}

// Run is route's
func Run() {

	r := gin.Default()

	r.POST("/receiver", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
	})

	r.Run(":14414")
}
