package service

import (
	"github.com/Qianjiachen55/gin-fanyi-project/global"
	"github.com/Qianjiachen55/gin-fanyi-project/initalize"
	"github.com/Qianjiachen55/gin-fanyi-project/model"
	"time"
)



func InsertDict(from string,to string) (*model.Dict,error) {
	dict := model.NewDict(from,to,time.Now().Format("2006-01-02 15:04:05"),0)

	result := global.GFP_MYSQL.Create(&dict)
	if result.Error != nil {
		return nil, result.Error
	}
	return dict,nil
}

func UpdateCount(id uint,count uint) error {
	db := initalize.ConnectMysql()
	result := db.Model(&model.Dict{}).Where("id = ?",id).Update("count",count)
	return result.Error
}

func Query(label string) ([]model.Dict,error) {
	db := initalize.ConnectMysql()
	var dict []model.Dict
	result := db.Where("`from` = ? or `to` = ?",label,label).Find(&dict)
	return dict,result.Error
}

