package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthLiveliness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"meaasge": "pong"})
}
