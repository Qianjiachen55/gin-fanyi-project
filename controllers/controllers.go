package controllers

import (
	"github.com/Qianjiachen55/gin-fanyi-project/global"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Test(ctx *gin.Context)  {
	ip := viper.GetString("mysql.ip")
	global.GFP_LOGGER.Debug(ip)
	ctx.String(200,"hello world!!!!")
}