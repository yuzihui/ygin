package v1

import (
	"ecloudsystem/model"
	"ecloudsystem/pkg/cache"
	"ecloudsystem/pkg/code"
	"ecloudsystem/pkg/db"
	"ecloudsystem/pkg/resp"
	"github.com/gin-gonic/gin"
	"time"
)

func TestInfo(c *gin.Context)  {

	 orm := db.Client.GetDbW()
	 menu  := new(model.TestModel)
	 menu.Id = 2
	 orm.First(&menu)

	 redisClient := cache.Client.Redis

	if _, err := redisClient.Set("test", "test", 1000 * time.Second).Result(); err != nil {

	}
	type a struct {
		A string `json:"a"`
		B string `json:"b"`
	}
	testYu := new(a)
	testYu.A = "21323"
	testYu.B = "21323"

	resp.Json(c, code.ServerError, nil, "")
}


