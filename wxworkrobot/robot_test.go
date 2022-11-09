package wxworkrobot

import (
	"context"
	"os"
	"reflect"
	"testing"
)

const RobotKey = "c275aafa-45a1-4a9d-b06d-97f05d5a673b"

func TestRobot_UploadMedia(t *testing.T) {
	file1, err := os.Open("./README.md")
	if err != nil {
		t.Errorf("invalid file1")
		return
	}
	type fields struct {
		Key string
	}
	type args struct {
		ctx      context.Context
		fileName string
		file     *os.File
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "TestRobot_UploadMedia 001",
			fields: fields{Key: RobotKey},
			args: args{
				ctx:      context.WithValue(context.Background(), ContextTraceKey, "001"),
				fileName: "test_file.md",
				file:     file1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Key: tt.fields.Key,
			}
			r.UploadMedia(tt.args.ctx, tt.args.fileName, tt.args.file)
		})
	}
}

func TestRobot_Send(t *testing.T) {
	type fields struct {
		Key string
	}
	type args struct {
		ctx     context.Context
		msgType MsgType
		body    SendReqBody
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRsp *SendRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestRobot_Send File 001",
			fields: fields{
				Key: RobotKey,
			},
			args: args{
				ctx:     context.WithValue(context.Background(), ContextTraceKey, "TestRobot_Send-001"),
				msgType: "file",
				body: SendReqBody{
					MsgType: "file",
					File: SendFileReqBody{
						MediaID: "3KFKb3TD-GNfWVhQL0ddqZrcRMwfqSh1ahV2JNCx2GdKFTA-F94WriKgrcD_fDCtB",
					},
				},
			},
			wantRsp: &SendRsp{
				ErrorCode: 0,
				ErrorMsg:  "ok",
			},
			wantErr: false,
		},
		{
			name: "TestRobot_Send Text 001",
			fields: fields{
				Key: RobotKey,
			},
			args: args{
				ctx:     context.WithValue(context.Background(), ContextTraceKey, "TestRobot_Send-Text-001"),
				msgType: "text",
				body: SendReqBody{
					MsgType: "text",
					Text: SendTextReqBody{
						Content:             "这是一条文本信息",
						MentionedList:       []string{"windealli"},
						MentionedMobileList: nil,
					},
				},
			},
			wantRsp: &SendRsp{
				ErrorCode: 0,
				ErrorMsg:  "ok",
			},
			wantErr: false,
		},
		{
			name: "TestRobot_Send Text 001",
			fields: fields{
				Key: RobotKey,
			},
			args: args{
				ctx:     context.WithValue(context.Background(), ContextTraceKey, "TestRobot_Send-Text-001"),
				msgType: "markdown",
				body: SendReqBody{
					MsgType: "markdown",
					Markdown: SendMarkdownReqBody{
						Content: "这是一条文本信息\n # Hello",
					},
				},
			},
			wantRsp: &SendRsp{
				ErrorCode: 0,
				ErrorMsg:  "ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				Key: tt.fields.Key,
			}
			gotRsp, err := r.Send(tt.args.ctx, tt.args.msgType, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("Send() gotRsp = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
