package tools

import "strings"

const (
	TypeJPG   = "image/jpeg"
	TypePNG   = "image/png"
	TypeGIF   = "image/gif"
	TypeText  = "text/plain"
	TypeMp3   = "audio/mpeg"
	TypeMp4   = "video/mp4"
	TypeOther = "application/octet-stream"
)

// GetContentType 获取文件的contentType类型
func GetContentType(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return TypeOther
	}
	switch filename[index+1:] {
	case "png":
		return TypePNG
	case "jpg", "jpeg":
		return TypeJPG
	case "mp3":
		return TypeMp3
	case "mp4", "mkv":
		return TypeMp4
	case "txt", "lrc", "ass", "nfo":
		return TypeText
	default:
		return TypeOther
	}
}

// GetBucketFromContentType 获取存储的bucket
func GetBucketFromContentType(contentType string) string {
	switch contentType {
	case TypeJPG, TypePNG, TypeGIF:
		return "image"
	case TypeMp4:
		return "video"
	case TypeMp3:
		return "music"
	case TypeText:
		return "text"
	default:
		return "other"
	}
}
