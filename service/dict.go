package service

import (
	"context"
	"github.com/Qianjiachen55/gin-fanyi-project/global"
	"github.com/Qianjiachen55/gin-fanyi-project/model"
	"github.com/spf13/viper"
	"time"
)



func InsertDictMysql(from string,to string) (*model.Dict,error) {
	dict := model.NewDict(from,to,time.Now().Format("2006-01-02 15:04:05"),0)

	result := global.GFP_MYSQL.Create(&dict)
	if result.Error != nil {
		return nil, result.Error
	}
	return dict,nil
}

func UpdateCountMysql(id uint,count uint) error {
	//db := initalize.ConnectMysql()
	result := global.GFP_MYSQL.Model(&model.Dict{}).Where("id = ?",id).Update("count",count)
	return result.Error
}

func QueryMysql(label string) ([]model.Dict,error) {
	//db := initalize.ConnectMysql()
	var dict []model.Dict
	result := global.GFP_MYSQL.Where("`from` = ? or `to` = ?",label,label).Find(&dict)
	return dict,result.Error
}

func InsertDictRedis(from string,to string)  error{
	ctx:=context.Background()
	TTL := viper.GetInt("redis.TTL")
	err := global.GFP_REDIS.Set(ctx,from,to, time.Duration(TTL)*time.Second).Err()
	if err !=nil{
		return err
	}
	err = global.GFP_REDIS.Set(ctx, to, from, time.Duration(TTL)*time.Second).Err()
	return err
}

func QueryRedis(from string)  (string,error){
	ctx := context.Background()
	to,err := global.GFP_REDIS.Get(ctx,from).Result()
	return to,err
}
