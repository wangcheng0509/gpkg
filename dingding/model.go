package dingding

// DingReq 钉钉请求
type DingReq struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

// Markdown Markdown内容
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// DingRsp 钉钉返回
type DingRsp struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
}
