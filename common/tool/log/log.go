package log

import (
	"context"
	"encoding/json"
	"fmt"
)

func CtxLogInfo(ctx context.Context, format string, data ...interface{}) {
	fmt.Printf(format+"\n", data...)
}

func CtxLogWarn(ctx context.Context, format string, data ...interface{}) {
	fmt.Printf(format+"\n", data...)
}

func CtxLogError(ctx context.Context, format string, data ...interface{}) {
	fmt.Printf(format+"\n", data...)
}

func Info(format string, data ...interface{}) {
	fmt.Printf(format+"\n", data...)
}

func Warn(format string, data ...interface{}) {
	fmt.Printf(format+"\n", data...)
}

func Error(format string, data ...interface{}) {
	fmt.Printf(format+"\n", data...)
}

func Json(format string, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Printf(format+"\n", string(res))
}
