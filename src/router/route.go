package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "server work!",
		})
	})

	return router
}
