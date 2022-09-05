package moudle

import (
	"context"
	"github.com/gin-gonic/gin"
	"index.app/internal/server"
)

type DataServer struct {
	server *server.DataService
}

func (s DataServer) DownloadLinkAndUpload(ctx context.Context, url string, name string, header map[string]string) map[string]interface{} {
	id, err := s.server.DownloadLinkAndUpload(ctx, url, name, header)
	return map[string]interface{}{"id": id, "err": err}
}

func (s DataServer) UploadObjectFromFile(ctx *gin.Context, field string) map[string]interface{} {
	res := map[string]interface{}{"id": "", "err": nil}
	f, err := ctx.FormFile(field)
	if err != nil {
		res["err"] = err
		return res
	}
	id, err := s.server.UploadObjectFromFile(ctx, f)
	if err != nil {
		res["err"] = err
		return res
	}
	res["id"] = id
	return res
}

func (s DataServer) AddText(ctx context.Context, name string, content string) map[string]interface{} {
	id, err := s.server.AddText(ctx, name, content)
	return map[string]interface{}{"id": id, "err": err}
}

func (s DataServer) UpdateText(ctx context.Context, id string, name string, content string) error {
	return s.server.UpdateText(ctx, id, name, content)
}

func (s DataServer) GetText(ctx context.Context, id string) map[string]interface{} {
	name, content, err := s.server.GetText(ctx, id)
	return map[string]interface{}{"name": name, "content": content, "err": err}
}

func (s DataServer) DeleteObject(ctx context.Context, id []string) error {
	return s.server.DeleteObject(ctx, id)
}

func (s DataServer) GetDownloadLink(ctx context.Context, id string) map[string]interface{} {
	link, err := s.server.GetDownloadLink(ctx, id)
	return map[string]interface{}{"link": link, "err": err}
}
