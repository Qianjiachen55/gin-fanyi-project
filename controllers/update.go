package controllers

import (
	"github.com/Qianjiachen55/gin-fanyi-project/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

func GetUpdate(ctx *gin.Context)  {
	id := ctx.Param("id")
	go func() {
		for count:= 0;count<100;count++	{
			service.UpdateCountMysql(cast.ToUint(id), cast.ToUint(count))
			time.Sleep(time.Second)

		}
	}()
	ctx.String(http.StatusOK,"updating......")
}
