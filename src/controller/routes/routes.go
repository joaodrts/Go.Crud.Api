package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaodrts/Go.Crud.Api/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("users", controller.GetAll)
	r.GET("users/:userEmail", controller.GetByEmail)
	r.POST("users", controller.Create)
	r.PUT("users/:userEmail", controller.Update)
	r.DELETE("users/:userEmail", controller.Delete)

}
