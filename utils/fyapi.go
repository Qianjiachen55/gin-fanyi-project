package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

func make_md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}


func translate(word string) ([]byte,error) { //调用api进行翻译
	data := make(url.Values)
	data["q"] = []string{word}
	data["from"] = []string{"auto"}
	data["to"] = []string{"auto"}
	data["appid"] = []string{viper.GetString("baidu.appid")}
	data["salt"] = []string{"65"}
	sk := viper.GetString("baidu.sk")
	s := data["appid"][0] + word + "65" + sk //密匙
	sign := make_md5(s)
	data["sign"] = []string{sign}
	res, err := http.PostForm("http://api.fanyi.baidu.com/api/trans/vip/translate", data)
	if err != nil {
		return nil,err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	//str := string(body)
	return body,nil
}
func Trans(words string) (string,error) { //翻译函数
	body,err := translate(words)
	js, err := simplejson.NewJson(body)
	if err != nil {
		return "",err
	}
	dst := js.Get("trans_result").GetIndex(0).Get("dst").MustString()
	return dst,nil
}

