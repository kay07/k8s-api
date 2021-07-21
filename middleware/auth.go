package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

//func Token() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		context.Header("Authorization","Bearer "+config.Conf("token"))
//		context.Next()
//	}
//}

func Cors()gin.HandlerFunc  {
	 cors:=cors.New(cors.Config{
			AllowOrigins:[]string{"*"},
			AllowMethods:[]string{"PUT","GET","POST","DELETE","PATCH"},
			AllowHeaders:[]string{"Origin","Authorization","Content-Type"},
			ExposeHeaders:[]string{"Content-Type"},
			AllowCredentials:true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			MaxAge:24*time.Hour,
		})
	return cors
}