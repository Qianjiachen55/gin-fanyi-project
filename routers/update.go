package routers

import (
	"github.com/Qianjiachen55/gin-fanyi-project/controllers"
	"github.com/gin-gonic/gin"
)

func LoadUpdateRouter(engine *gin.Engine)  {
	engine.GET("/api/v1/update",controllers.GetUpdate)
}
