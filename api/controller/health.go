package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description Health Check to ensure API is working and returning a response
// @Summary Health Check to ensure API is working and returning a response
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/wallets/ping [get]
func HealthLiveliness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"meaasge": "pong"})
}
