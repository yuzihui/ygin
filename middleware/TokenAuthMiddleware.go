package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)


func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//token := c.Query("token")
		//if len(token) == 0 {
		//	a:= gin.H{"A":"token 验证不通过"}
		//	c.JSON(401, a)
		//	c.Abort()
		//}

		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
		c.Next()
	}
}
