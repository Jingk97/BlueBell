package routers

import (
	"BlueBell/controllers"
	"github.com/gin-gonic/gin"
)

func InitApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", controllers.ApiController{}.ApiUse)

	}
}
