package api

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"k8s-api/middleware"
	"net/http"
	"strconv"
	"sync"
)

var client *http.Client
var once sync.Once
func Client() *http.Client{
	once.Do(func() {
		client=&http.Client{
			Transport: &http.Transport{
				TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
			},
		}
	})
	return client
}
// @Summary 获取所有deployment
// @Description 详情
// @Tags k8s-deployment
// @Param page query string false "page"
// @Success 200
// @Failure 400
// @Router /alldep [get]
func GetDep(c *gin.Context)  {
	//page:=c.Param("page")
	//if page=="/"{
	//	page="1"
	//}

	page,bool:=c.GetQuery("page")
	if bool==false{
		page=`1`
	}
	con:=GetConDescription()
	//分页
	num,_:=strconv.Atoi(page)
	results:=middleware.Page(num,con)

	c.JSON(200,results)

}

