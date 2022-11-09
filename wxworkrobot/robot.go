package wxworkrobot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type Robot struct {
	Key string
}

func NewRobot(ctx context.Context, key string) (robot *Robot, err error) {
	robot = &Robot{Key: key}
	return robot, err
}

// UploadMedia : 文件上传接口
func (r *Robot) UploadMedia(ctx context.Context, fileName string, file *os.File) (rsp *UploadMediaRsp, err error) {
	LogInfoContextf(ctx, "UploadMedia")

	// 创建buffer 和 buffer writer
	buf := &bytes.Buffer{}
	bw := multipart.NewWriter(buf)

	fw, err := bw.CreateFormFile("media", fileName)
	if err != nil {
		LogErrorContextf(ctx, "CreateFormFile error, %+v", err)
		return rsp, err
	}
	n, err := io.Copy(fw, file)
	if err != nil {
		LogErrorContextf(ctx, "io.Copy error, %+v", err)
		return rsp, err
	}
	LogDebugContextf(ctx, "io.Copy() , %d", n)

	bw.Close() // 此处不可以使用defer，否则会导致buf中没有内容

	uploadURL := fmt.Sprintf(URLFormatQYAPIUploadMedia, r.Key)
	req, err := http.NewRequest(http.MethodPost, uploadURL, buf)
	if err != nil {
		LogErrorContextf(ctx, "http.NewRequest error, %+v", err)
		return rsp, err
	}
	//
	contentType := bw.FormDataContentType()
	req.Header.Add("Content-Type", contentType)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return rsp, err
	}
	//
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogErrorContextf(ctx, "ioutil.ReadAll error, %+v", err)
		return
	}
	resp.Body.Close()
	LogInfoContextf(ctx, "UploadMedia: %s", string(body))

	rsp = &UploadMediaRsp{}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		LogErrorContextf(ctx, "json.UnMarshal error, %+v", err)
		return rsp, err
	}

	if rsp.ErrorCode != 0 {
		LogErrorContextf(ctx, "UploadMedia error, %d, %s", rsp.ErrorCode, rsp.ErrorMsg)
		return rsp, fmt.Errorf("UploadMedia error, %d, %s", rsp.ErrorCode, rsp.ErrorMsg)
	}

	LogDebugContextf(ctx, "UploadMedia success, %+v", rsp)
	return
}

// Send : 发送消息
func (r *Robot) Send(ctx context.Context, msgType MsgType, body SendReqBody) (rsp *SendRsp, err error) {
	LogInfoContextf(ctx, "Send, %s, %+v", msgType, body)

	reqBody, err := json.Marshal(body)
	if err != nil {
		LogErrorContextf(ctx, "json.Marshal error, %+v", err)
		return
	}

	url := fmt.Sprintf(URLFormatQYAPISend, r.Key)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		LogErrorContextf(ctx, "http.NewRequest error, %+v", err)
		return
	}

	// 发送POST请求
	cli := http.Client{}
	httpRsp, err := cli.Do(req)
	if err != nil {
		LogErrorContextf(ctx, "http client.Do error, %+v", err)
		return
	}
	defer httpRsp.Body.Close()

	bodyData, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		LogErrorContextf(ctx, "ioutil.ReadAll error, %+v", err)
		return
	}
	LogInfoContextf(ctx, "ioutil.ReadAll success, %s", string(bodyData))

	rsp = &SendRsp{}
	err = json.Unmarshal(bodyData, rsp)
	if err != nil {
		LogErrorContextf(ctx, "json unmarshal error, %+v", err)
		return
	}

	LogDebugContextf(ctx, "Send success, %+v", rsp)
	return
}
