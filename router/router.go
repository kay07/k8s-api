package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"k8s-api/api"
	_ "k8s-api/docs"
	"k8s-api/middleware"
)

func Init()*gin.Engine  {
	router:=gin.Default()

	router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//使用自定义token

	router.Use(middleware.Cors())
	//get,post
	router.GET("/alldep",api.GetDep)

	router.POST("/svc",api.Create)
	router.POST("/svctest",api.Creat)
	return router
}
