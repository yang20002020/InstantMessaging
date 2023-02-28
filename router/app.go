package router

import (
	"InstantMessaging/docs"
	"InstantMessaging/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r

}
