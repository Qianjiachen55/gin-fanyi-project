package service

import (
	"fmt"
	"github.com/Qianjiachen55/gin-fanyi-project/global"
	"github.com/Qianjiachen55/gin-fanyi-project/initalize"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInsertDict(t *testing.T) {
	dict, err := InsertDictMysql("apple", "苹果")
	fmt.Println(err)
	fmt.Println(dict)
}
func TestUpdateCount(t *testing.T) {
	if err := UpdateCountMysql(2, 1); err != nil {
		fmt.Println("error:",err)
	}
}
func TestQuery(t *testing.T) {
	dictList,err := QueryMysql("苹果")
	if err != nil{
		fmt.Println(err)
	}
	for _,dict := range dictList{
		fmt.Println(dict)
	}
}
func TestInsertDictRedis(t *testing.T) {
	if err := InsertDictRedis("苹果","apple");err != nil{
		fmt.Println(err)
	}
}

func init() {

	// Find home directory.
	//home, err := os.UserHomeDir()
	home := "../conf"
	//cobra.CheckErr(err)

	// Search config in home directory with name ".gin-fanyi-project" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("conf")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	global.GFP_MYSQL = initalize.ConnectMysql()

	global.GFP_REDIS = initalize.RedisClient()

	initalize.InitMysql(global.GFP_MYSQL)
	initalize.InitRedis(global.GFP_REDIS)
}
