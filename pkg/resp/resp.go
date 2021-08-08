package resp

import (
	errno "ecloudsystem/pkg/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    		`json:"code"`    // 业务码
	Message string 		`json:"message"` // 描述信息
	Data    interface{} `json:"data"`    // 数据
}


func Json(ctx *gin.Context, code int, data interface{},  message string)  {

	var res Response
	res.Code = code
	res.Data = data

	if len(message) > 0{
		res.Message = message
	} else {
		res.Message = errno.Text(code)
	}
	ctx.JSON(http.StatusOK, res)
	return
}