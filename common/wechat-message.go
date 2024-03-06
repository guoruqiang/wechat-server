package common

import "encoding/xml"

type WeChatMessageRequest struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	Event        string   `xml:"Event"`
	EventKey     string   `xml:"EventKey"`
	MsgId        int64    `xml:"MsgId"`
	MsgDataId    int64    `xml:"MsgDataId"`
	Idx          int64    `xml:"Idx"`
}

type WeChatMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

func ProcessWeChatMessage(req *WeChatMessageRequest, res *WeChatMessageResponse) {
	if req.MsgType == "event" && req.Event == "CLICK" {
		switch req.EventKey {
		case "USER_VERIFICATION":
			code := GenerateAllNumberVerificationCode(6)
			RegisterWeChatCodeAndID(code, req.FromUserName)
			res.Content = code
		}
	} else {
		switch req.Content {
		case "验证码":
			code := GenerateAllNumberVerificationCode(6)
			RegisterWeChatCodeAndID(code, req.FromUserName)
			res.Content = code
		}
	}
}
