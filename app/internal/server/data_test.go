package server

import (
	"bytes"
	"context"
	"fmt"
	"testing"
)

func TestDataService_GetDownloadLink(t *testing.T) {
	fmt.Println(data1.GetDownloadLink(context.Background(), "62caa90f45355445b731a6bd"))
}

func TestDataService_UploadObjectFromFile(t *testing.T) {
	data := []byte("hello word")
	id, err := data1.UploadFromReader(context.Background(), bytes.NewReader(data), "hello.txt", int64(len(data)))
	fmt.Println(id)
	fmt.Println(err)
}
