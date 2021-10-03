package service

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInsertDict(t *testing.T) {
	dict, err := InsertDict("apple", "苹果")
	fmt.Println(err)
	fmt.Println(dict)
}
func TestUpdateCount(t *testing.T) {
	if err := UpdateCount(2, 1); err != nil {
		fmt.Println("error:",err)
	}
}
func TestQuery(t *testing.T) {
	dictList,err := Query("苹果")
	if err != nil{
		fmt.Println(err)
	}
	for _,dict := range dictList{
		fmt.Println(dict)
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
}
