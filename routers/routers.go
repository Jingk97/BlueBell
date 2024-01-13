package routers

import (
	"BlueBell/logger"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	InitFrontRoters(r)
	InitAdminRouters(r)
	InitApiRouters(r)
	return r
}
