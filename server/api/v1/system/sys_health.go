package system

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
)

type HealthApi struct{}

func (h HealthApi) Health(c *gin.Context) {
	now := time.Now()
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"msg":     "server is healthy",
		"ip":      c.ClientIP(),
		"version": global.Version,
		"mode":    gin.Mode(),
		"uptime":  fmt.Sprintf("%vs", now.Sub(global.GVA_START_TIME).Seconds()),
	})
}
