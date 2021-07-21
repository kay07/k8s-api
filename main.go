package main

import (
	"github.com/gin-gonic/gin"
	"k8s-api/router"
)

// @Title kubernetes API
// @version 1.0
func main()  {
	gin.SetMode(gin.ReleaseMode)
	r:=router.Init()

	r.Run()
}
