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
	if req.MsgType == "event" {
		switch req.Event {
		case "CLICK":
			switch req.EventKey {
			case "USER_VERIFICATION":
				code := GenerateAllNumberVerificationCode(6)
				RegisterWeChatCodeAndID(code, req.FromUserName)
				res.Content = code
			}
		case "subscribe": // 处理关注事件
			res.Content = "欢迎关注小锦亿。使用西农er's ChatGPT的小伙伴点击下方的获取验证码即可登录哟，再也不用担心忘记用户名了。有什么问题联系管理员。" // 自动回复的消息内容
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
