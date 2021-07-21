package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
// @Summary 创建svc
// @Description 详情
// @Tags k8s-svc
// @Param name query string true "name"
// @Param port query string true "port"
// @Param replicas query string true "replicas"
// @Param image query string true "image"
// @Param env query string false "env"
// @Param volumes query string false "volumes"
// @Success 200
// @Failure 400
// @Router /svc [post]
func Create(c *gin.Context)  {
	name,_:=c.GetQuery("name")
	port,_:=c.GetQuery("port")
	replicas,_:=c.GetQuery("replicas")
	image,_:=c.GetQuery("image")
	env,_:=c.GetQuery("env")
	volumes,_:=c.GetQuery("volumes")

	//创建deployment
	err:=CreateDep(name,replicas,image,env,volumes)

	//创建service
	err=CreateSvc(name,port)
	fmt.Println(err)
	c.JSON(200,"ok")
}
