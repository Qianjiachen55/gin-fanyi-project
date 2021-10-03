package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Test(ctx *gin.Context)  {
	ip := viper.GetString("mysql.ip")
	fmt.Println(ip)
	ctx.String(200,"hello world!!!!")
}