package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"telcobss/internal/common/models"
	"telcobss/internal/order/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.OrderRequest
	if err := c.BindJSON(&req); err != nil {
		logrus.Errorf("invalid order request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.OrderID == "" || req.CustomerID == "" || req.ProductCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		return
	}

	if err := h.service.PlaceOrder(c.Request.Context(), &req); err != nil {
		logrus.Errorf("failed place order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to place order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order_id": req.OrderID, "status": "placed"})
}
