package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hopers-backend/data"
	"time"
)

// @BasePath /api

// PingExample godoc
// @Summary health
// @Schemes
// @Description ensure api is responding
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Health
// @Router /health [get]
func Health(c *gin.Context) {
	Health := data.Health{}
	Health.Timestamp = time.Now().Format(time.RFC3339)
	Health.Status = "In hopers-backend We Trust"
	HealthJson, _ := json.Marshal(Health)
	c.String(200, string(HealthJson))
}
