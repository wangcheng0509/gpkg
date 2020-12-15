package exception

// SendDingdingReq 发送钉钉请求参数
type sendDingdingReq struct {
	DingReq dingReq `json:"content"`
	Webhook string  `json:"webhook"`
	Secret  string  `json:"secret"`
}

// DingReq 钉钉请求
type dingReq struct {
	Msgtype  string   `json:"msgtype"`
	Markdown markdown `json:"markdown"`
}

// Markdown Markdown内容
type markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
