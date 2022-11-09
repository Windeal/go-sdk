package wxworkrobot

type MsgType string

const (
	MsgTypeFile MsgType = "file"
)

type UploadMediaRsp struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
	Type      string `json:"file,omitempty"`
	MediaID   string `json:"media_id,omitempty"`
	CreateAt  string `json:"created_at,omitempty"`
}

type SendReqBody struct {
	MsgType  MsgType             `json:"msgtype"`
	File     SendFileReqBody     `json:"file,omitempty"`
	Text     SendTextReqBody     `json:"text,omitempty"`
	Markdown SendMarkdownReqBody `json:"markdown,omitempty"`
}

type SendFileReqBody struct {
	MediaID string `json:"media_id,omitempty"`
}

type SendTextReqBody struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type SendMarkdownReqBody struct {
	Content string `json:"content"`
}

type SendRsp struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}
