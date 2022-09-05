package tools

import (
	"errors"
	"fmt"
)

// ReturnError 返回错误信息
func ReturnError(format string, a ...any) error {
	return errors.New(fmt.Sprintf(format, a...))
}
