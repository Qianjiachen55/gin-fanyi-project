// Package cmd /*
package cmd

import (
	"github.com/Qianjiachen55/gin-fanyi-project/global"
	"github.com/Qianjiachen55/gin-fanyi-project/initalize"
	"github.com/Qianjiachen55/gin-fanyi-project/routers"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start gfp server",
	Long: `use go run main.go start to start server or go build -o app && ./app start to start server`,
	Run: Run,
}

func Run(cmd *cobra.Command, args []string) {
	engine := gin.New()
	global.GFP_MYSQL = initalize.ConnectMysql()
	global.GFP_REDIS = initalize.RedisClient()
	global.GFP_LOGGER = initalize.Logger()
	initalize.InitMysql(global.GFP_MYSQL)
	initalize.InitRedis(global.GFP_REDIS)
	engine.Use(ginzap.Ginzap(global.GFP_LOGGER, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(global.GFP_LOGGER, true))
	routers.LoadRootRouter(engine)
	routers.LoadDictRouter(engine)
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
}
