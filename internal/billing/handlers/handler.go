package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"telcobss/internal/billing/service"
)

type BillingHandler struct {
	service *service.BillingService
}

func NewBillingHandler(s *service.BillingService) *BillingHandler {
	return &BillingHandler{service: s}
}

func (h *BillingHandler) GetInvoice(c *gin.Context) {
	invoiceID := c.Param("id")
	if invoiceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invoice id required"})
		return
	}

	invoice, err := h.service.FetchInvoice(c.Request.Context(), invoiceID)
	if err != nil {
		logrus.Errorf("failed fetch invoice: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}
