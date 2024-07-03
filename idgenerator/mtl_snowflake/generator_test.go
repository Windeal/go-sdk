package mtl_snowflake

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	// type fields struct {
	// 	mutex            *sync.Mutex
	// 	config           *Config
	// 	preSets          *PreSets
	// 	timelineProgress []int64
	// 	curTimeline      int64
	// 	seq              uint64
	// }
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		// fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "生成10W个ID",
			// fields: fields{},
			args: args{
				ctx: context.Background(),
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// gen := &Generator{
			// 	mutex:            tt.fields.mutex,
			// 	config:           tt.fields.config,
			// 	preSets:          tt.fields.preSets,
			// 	timelineProgress: tt.fields.timelineProgress,
			// 	curTimeline:      tt.fields.curTimeline,
			// 	seq:              tt.fields.seq,
			// }

			gen, err := NewGenerator(context.Background())
			if err != nil {
				fmt.Println("初始化生成器实例:", err)
				return
			}

			// 生成一个文件用于写入id
			file, err := os.OpenFile("./id.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("无法打开文件:", err)
				return
			}
			defer file.Close()
			writer := bufio.NewWriter(file)

			id1 := int64(0)
			id2 := int64(1)

			for i := 0; i < 100*1000; i++ {
				got, err := gen.Generate(tt.args.ctx)
				if (err != nil) != tt.wantErr {
					t.Errorf("Generator.Generate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				id1 = id2
				id2 = got
				if id1 >= id2 {
					t.Errorf("invalid id %d, %d", id1, id2)
				}

				// if got != tt.want {
				// 	t.Errorf("Generator.Generate() = %v, want %v", got, tt.want)
				// }
				writer.WriteString(strconv.FormatInt(got, 10) + "\n")
			}

			writer.Flush()
		})
	}
}
