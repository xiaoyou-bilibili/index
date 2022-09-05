package service

import (
	"context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	s := CreateKvRocksService()
	err := s.SetText(context.Background(), "123", []byte("hello word"))
	fmt.Println(err)
	a, err := s.GetTextByte(context.Background(), "123")
	fmt.Println(err)
	fmt.Println(string(a))
}
