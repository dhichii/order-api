package router

import (
	"net/http"
	"order-api/src/adapter"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	handler := adapter.Init()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "server work!",
		})
	})

	order := router.Group("/orders")
	order.POST("", handler.Order.CreateOrderHandler)
	order.GET("", handler.Order.GetOrdersHandler)
	order.PUT("/:orderId", handler.Order.UpdateOrderHandler)
	order.DELETE("/:orderId", handler.Order.DeleteOrderHandler)

	return router
}
