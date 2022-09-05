package tools

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetContentType(t *testing.T) {
	//assert.Same(t, TypeJPG, GetContentType("21323.jpg"))
}

func TestName(t *testing.T) {
	file := "q.aaa.222"
	index := strings.LastIndex(file, ".")
	fmt.Println(file[:index])
}
