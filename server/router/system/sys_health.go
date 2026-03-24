package system

import (
	"github.com/gin-gonic/gin"
)

type HealthRouter struct {
}

func (h HealthRouter) InitHealthRouter(Router *gin.RouterGroup) {
	Router.GET("health", healthApi.Health) // 健康检查
}
