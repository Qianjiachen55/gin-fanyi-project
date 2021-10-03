package initalize

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql()  *gorm.DB{
	ip := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	passwd := viper.GetString("mysql.passwd")
	database := viper.GetString("mysql.database")
	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		passwd,
		ip,
		port,
		database,
		)
	fmt.Println(dsn)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		fmt.Println(err)
		panic(err)
	}


	return db
}