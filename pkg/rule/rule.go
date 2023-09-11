package rule

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/zxfishhack/mirror/pkg/model"
	"github.com/zxfishhack/mirror/pkg/storage"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

type rule struct {
	model.Rule
	storage storage.IStorage
}

type RuleController struct {
	Ctx  iris.Context
	Data *rule
}

func newRule(r model.Rule, create storage.CreateStorageFunc) (c *RuleController, err error) {
	s, err := create(r.Prefix)
	if err != nil {
		return
	}
	c = &RuleController{
		Data: &rule{
			Rule:    r,
			storage: s,
		},
	}

	return
}

func (r *RuleController) GetByWildcard(p string) (err error) {
	lp := path.Join(r.Data.Prefix, p)
	b, err := r.Data.storage.Get(lp)
	for errors.Is(err, os.ErrNotExist) {
		var u *url.URL
		u, err = url.Parse(r.Data.Upstream)
		if err != nil {
			break
		}
		u.Path = path.Join(u.Path, r.Data.ReplacePrefixWith, p)
		var resp *http.Response
		resp, err = http.Get(u.String())
		if err != nil {
			break
		}
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			break
		}
		if resp.StatusCode != http.StatusOK {
			r.Ctx.StatusCode(resp.StatusCode)
			r.Ctx.ContentType(resp.Header.Get("Content-Type"))
			break
		}
		err = r.Data.storage.Put(lp, b)
		if err != nil {
			err = nil
			break
		}
	}
	if err == nil {
		if r.Ctx.ResponseWriter().StatusCode() == http.StatusOK {
			r.Ctx.ContentType(mime.TypeByExtension(filepath.Ext(lp)))
		}
		_, err = r.Ctx.Write(b)
		return err
	}
	return nil
}
