package plug

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := c.PostForm("cmd")
		param := c.PostForm("param")

		if cmd == "" || param == "" {

			c.JSON(http.StatusOK, gin.H{
				"msg": "Cmd or param not found",
			})

			c.Abort()

		}
	}
}

func VerifyRouteList() gin.HandlerFunc {
	return func(c *gin.Context) {

		cmd := c.PostForm("cmd")

		if _, ok := RouteList[cmd]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Cmd not found",
			})

			c.Abort()
		}
	}
}
