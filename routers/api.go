package routers

import (
	"ecloudsystem/controller/api/v1"
	"ecloudsystem/middleware"
	"github.com/gin-gonic/gin"
)

func setApiRoute(r *gin.Engine) {

	apiRoute := r.Group("/api")
	{

		v1Route := apiRoute.Group("v1")
		{
			/**
			 * v1 版本需要中间件版本
			 */
			auth:= v1Route.Use(middleware.TokenAuthMiddleware())
			{
				auth.GET("/test", v1.TestInfo)
				//auth.GET("/test10", v1.TestInfo)
			}
		}
	}

}



