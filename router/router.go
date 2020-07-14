package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhanxiaox/roy/plug"
)

func Setup() {
	plug.Setup()
}

// Run is init function of application
func Run() {

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.Use(plug.VerifyPost(), plug.VerifyRouteList())

		v1.POST("/receiver", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})

	r.Run(":14414")
}
