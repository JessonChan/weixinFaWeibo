package controllers

import (
	"../models"
	"encoding/xml"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type WeixinController struct {
	beego.Controller
}

func (this *WeixinController) Get() {
	//TODO check
	this.Ctx.WriteString(this.Input().Get("echostr"))
	println("checked" + this.Input().Get("nonce"))
}

func (this *WeixinController) Post() {
	body, e := ioutil.ReadAll(this.Ctx.Request.Body)
	if e != nil {
		println(e.Error())
	}
	textMsg := &models.TextMsg{}
	e = xml.Unmarshal(body, textMsg)
	if e != nil {
		println("xml Unmarshal")
		println(e.Error())
	}
	println(textMsg.Content)

	body, e = xml.Marshal(responseText(textMsg))
	if e != nil {
		println("xml Marshal")
		println(e.Error())
	}

	println(string(body))
	this.Ctx.WriteString(string(body))

}

func responseText(textMsg *models.TextMsg) (textResponse *models.TextResponse) {
	textResponse = &models.TextResponse{}
	textResponse.ToUserName = textMsg.FromUserName
	textResponse.FromUserName = textMsg.ToUserName
	textResponse.CreateTime = int(time.Now().Unix())
	textResponse.MsgType = "text"
	textResponse.FuncFlag = 0

	msg := textMsg.Content
	isSuccess := false
	if strings.HasPrefix(msg, "#") {
		isSuccess = faweibo(textMsg.FromUserName, msg[1:])
	}
	if !isSuccess {
		textResponse.Content = "hi,请点击下面的链接来绑定账号 http://106.3.46.54/weibo?weixinName=" + textMsg.FromUserName
	}else{
		textResponse.Content="send ok"
	}
	return textResponse
}

func faweibo(name, content string) bool {
	urlValues := make(url.Values)
	urlValues.Set("access_token", models.GetToken(name))
	urlValues.Set("status", content)
	_, err := http.Post("https://api.weibo.com/2/statuses/update.json", "application/x-www-form-urlencoded", strings.NewReader(urlValues.Encode()))
	if err != nil {
		return false
	}
	return true
}
