package routers

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) *gin.Engine {

	 setApiRoute(r)

	return  r
}