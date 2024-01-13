package routers

import (
	"BlueBell/controllers"
	"github.com/gin-gonic/gin"
)

func InitAdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", controllers.AdminController{}.AdminList)
	}
}
