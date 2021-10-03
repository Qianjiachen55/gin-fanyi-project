package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func Test(t *testing.T)  {
	ret,_ := Trans("苹果")
	fmt.Println(ret)
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