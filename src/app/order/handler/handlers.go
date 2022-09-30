package handler

import (
	"net/http"
	"order-api/src/app/order"
	"order-api/src/app/order/handler/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services order.Service
}

// CreateOrderHandler handler to create new order
func (h *Handler) CreateOrderHandler(c *gin.Context) {
	request := request.Request{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := h.services.CreateOrder(request.MapToRecord()); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]string{
		"message": "order created successfully",
	})
}

// GetOrdersHandler handler to get list of orders
func (h *Handler) GetOrdersHandler(c *gin.Context) {
	orders, err := h.services.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// UpdateOrderHandler handler to update the existing order
func (h *Handler) UpdateOrderHandler(c *gin.Context) {
	request := request.Request{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	id, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := h.services.UpdateOrder(id, request.MapToRecord()); err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, map[string]string{"message": "order not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "order has been updated successfully",
	})
}

// DeleteOrderHandler handler to delete the order data
func (h *Handler) DeleteOrderHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := h.services.DeleteOrder(id); err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, map[string]string{"message": "order not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "order has been deleted successfully",
	})
}

func NewHandler(service order.Service) *Handler {
	return &Handler{service}
}
