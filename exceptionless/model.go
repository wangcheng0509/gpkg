package exceptionless

type exLessReq struct {
	Type    int    `json:"type"`    // 1 log，2 error
	APIKey  string `json:"apiKey"`  // apiKey
	Message string `json:"message"` // message
}

type exLessRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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
