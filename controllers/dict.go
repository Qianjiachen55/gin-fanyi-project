package controllers

import (
	"github.com/Qianjiachen55/gin-fanyi-project/model"
	"github.com/Qianjiachen55/gin-fanyi-project/service"
	"github.com/Qianjiachen55/gin-fanyi-project/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostDict(ctx *gin.Context) {
	from := ctx.PostForm("from")
	to, err := utils.Trans(from)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, model.Response{
			Code:    500,
			Message: err.Error(),
		})
	}

	dict, err := service.InsertDictMysql(from, to)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, model.Response{
			Code:    500,
			Message: err.Error(),
		})
	}
	err = service.InsertDictRedis(from,to)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, model.Response{
			Code:    500,
			Message: err.Error(),
		})
	}
	resFrom := dict.From
	resTo := dict.To
	if from != dict.From {
		resTo, resFrom = resFrom, resTo
	}
	ctx.IndentedJSON(http.StatusOK, model.Response{
		Code:    200,
		Message: "OK",
		Data: model.Data{
			From: resFrom,
			To:   resTo,
		},
	})
}
