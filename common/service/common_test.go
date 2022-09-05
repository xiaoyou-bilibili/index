package service

import (
	"context"
	"fmt"
	"testing"
)

func TestGetStringValue(t *testing.T) {
	InitConsul()
	fmt.Println(GetStringValue(context.Background(), "index/music/qq_api_cookie"))
}
