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
	MsgType MsgType         `json:"msgtype"`
	File    SendFileReqBody `json:"file,omitempty"`
}

type SendFileReqBody struct {
	MediaID string `json:"media_id,omitempty"`
}

type SendRsp struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}
