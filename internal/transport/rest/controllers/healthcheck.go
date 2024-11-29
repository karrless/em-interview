package controllers

import "github.com/gin-gonic/gin"

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// @Summary Health check
// @Tags Healthcheck
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (hc *HealthCheckController) GetHealthCheck(c *gin.Context) {
	c.JSON(200, "pong")
}
