package idgenerator

import (
	"context"
	"github.com/windeal/go-sdk/idgenerator/logger"
	"github.com/windeal/go-sdk/idgenerator/mtl_snowflake"
)

var defaultGenerate *mtl_snowflake.Generator

func init() {
	var err error
	ctx := context.Background()
	defaultGenerate, err = mtl_snowflake.NewGenerator(ctx)
	if err != nil {
		logger.LogFatalContextf(ctx, "mtl_snowflake.NewGenerator error, %+v", err)
	}
}

func GenerateID(ctx context.Context) (int64, error) {
	return defaultGenerate.Generate(ctx)
}
