package routers

import (
	"github.com/Qianjiachen55/gin-fanyi-project/controllers"
	"github.com/gin-gonic/gin"
)

func LoadDictRouter(engine *gin.Engine)  {
	engine.POST("/api/v1/dict",controllers.PostDict)
	//engine.GET("/api/v1/dict")
}
