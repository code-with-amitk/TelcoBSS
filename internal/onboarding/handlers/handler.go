package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"telcobss/internal/common/models"
	"telcobss/internal/onboarding/service"
)

type OnboardingHandler struct {
	service *service.OnboardingService
}

func NewOnboardingHandler(s *service.OnboardingService) *OnboardingHandler {
	return &OnboardingHandler{service: s}
}

func (h *OnboardingHandler) CreateCustomer(c *gin.Context) {
	var req models.Customer
	if err := c.BindJSON(&req); err != nil {
		logrus.Errorf("invalid onboarding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.CustomerID == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		return
	}

	if err := h.service.CreateCustomer(c.Request.Context(), &req); err != nil {
		logrus.Errorf("failed create customer: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to onboard customer"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"customer_id": req.CustomerID, "status": "onboarded"})
}
