package routers

import (
	"BlueBell/controllers"
	"github.com/gin-gonic/gin"
)

func InitFrontRoters(r *gin.Engine) {
	frontRouters := r.Group("/")
	{
		frontRouters.GET("/", controllers.FrontController{}.FrontUse)
	}
}
