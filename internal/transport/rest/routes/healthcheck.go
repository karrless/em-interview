package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karrless/em-interview/internal/transport/rest/controllers"
)

func HealthCheckRoutes(r *gin.RouterGroup) {
	healthcheckController := controllers.NewHealthCheckController()
	r.GET("/ping", healthcheckController.GetHealthCheck)
}
