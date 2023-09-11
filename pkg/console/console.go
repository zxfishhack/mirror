package console

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"mime"
	"path/filepath"
	"strings"
)

type Handler struct {
	Ctx iris.Context
}

func (c *Handler) GetByWildcard(path string) error {
	if path == "" || path == "/" {
		path = "index.html"
	}
	if b, err := Asset(path); err == nil {
		if strings.HasSuffix(path, "html") {
			c.Ctx.Header("Cache-Control", "public, max-age=604800, immutable")
		}
		c.Ctx.ServeFile(path, false)
		c.Ctx.ContentType(mime.TypeByExtension(filepath.Ext(path)))
		c.Ctx.Write(b)
	}
	return nil
}

func Handle(app *mvc.Application) {

}
