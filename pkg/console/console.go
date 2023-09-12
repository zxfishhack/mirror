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

func (c *Handler) GetByWildcard(path string) (err error) {
	if path == "" || path == "/" {
		path = "index.html"
	}
	var b []byte
	if b, err = Asset(path); err == nil {
		mt := mime.TypeByExtension(filepath.Ext(path))
		if strings.Contains(mt, "text/html") {
			c.Ctx.Header("Cache-Control", "no-cache, max-age=0")
		} else {
			c.Ctx.Header("Cache-Control", "public, max-age=604800, immutable")
		}
		// c.Ctx.ServeFile(path, false)
		c.Ctx.ContentType(mt)
		if c.Ctx.ClientSupportsGzip() {
			_, err = c.Ctx.WriteGzip(b)
		} else {
			_, err = c.Ctx.Write(b)
		}
	}
	return
}

func Handle(app *mvc.Application) {
	app.Handle(&Handler{})
}
