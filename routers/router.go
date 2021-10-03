package routers

import (
	"github.com/Qianjiachen55/gin-fanyi-project/controllers"
	"github.com/gin-gonic/gin"
)

func LoadRouter(engine *gin.Engine)  {
	engine.GET("/",controllers.Test)
}