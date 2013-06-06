package models

import "encoding/xml"

type TextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        int64
}

type TextResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	FuncFlag     int
}
