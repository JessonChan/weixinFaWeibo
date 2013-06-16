package controllers

import (
	"../models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type WeiboController struct {
	beego.Controller
}

func (this *WeiboController) Get() {
	weixinName := this.Input().Get("weixinName")
	code := this.Input().Get("code")
	println(weixinName)
	println(code)

	if weixinName != "" {
		this.SetSession("weixinName", weixinName)
		this.Data["url"] = "https://api.weibo.com/oauth2/authorize?client_id=1917201163&response_type=code&redirect_uri=http://106.3.46.54/weibo"
		this.Data["data"] = "点击这里 绑定"
	}

	if code != "" {
		userName, _ := this.GetSession("weixinName").(string)
		println("name")
		println(userName)

		token := getAccessToken(code)
		user := &models.User{}
		user.WeixinUserName = userName
		user.Token = token
		e := models.InsertUser(user)
		if e != nil {
			println(e.Error())
		}
		this.Data["tips"] = "绑定成功,返回后直接给他发信息就能自动发到微博了，目前只支持发文本的消息"
	}
	this.TplNames = "weibo.tpl"

}
func getAccessToken(code string) string {
	urlValues := make(url.Values)
	urlValues.Set("client_id", "1917201163")
	urlValues.Set("client_secret", "b271c66875dc6c0e4d076a0c017189fd")
	urlValues.Set("grant_type", "authorization_code")
	urlValues.Set("code", code)
	urlValues.Set("redirect_uri", "http://106.3.46.54/weibo")
	r, _ := http.Post("https://api.weibo.com/oauth2/access_token", "application/x-www-form-urlencoded", strings.NewReader(urlValues.Encode()))
	body, _ := ioutil.ReadAll(r.Body)
	println(string(body))
	defer r.Body.Close()
	a := new(models.AccessToken)
	json.Unmarshal(body, &a)
	token := a.Access_token
	return token
}

