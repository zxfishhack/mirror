package rule

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/zxfishhack/mirror/pkg/model"
	"github.com/zxfishhack/mirror/pkg/storage"
	"github.com/zxfishhack/mirror/pkg/utils"
	"io"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type rule struct {
	model.Rule
	storage storage.IStorage
	clt     *http.Client
}

type RuleController struct {
	Ctx  iris.Context
	Data *rule
}

func newRule(r model.Rule, create storage.CreateStorageFunc) (c *RuleController, err error) {
	s, err := create(utils.String(r.Prefix))
	if err != nil {
		return
	}
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: defaultTransportDialContext(&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}),
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	if utils.String(r.ProxyUrl) != "" {
		if u, e := url.Parse(utils.String(r.ProxyUrl)); e == nil {
			transport.Proxy = http.ProxyURL(u)
		}
	}
	c = &RuleController{
		Data: &rule{
			Rule:    r,
			storage: s,
			clt:     &http.Client{Transport: transport},
		},
	}

	return
}

func (r *RuleController) GetByWildcard(p string) (err error) {
	lp := path.Join(utils.String(r.Data.Prefix), p)
	b, err := r.Data.storage.Get(lp)
	for errors.Is(err, os.ErrNotExist) {
		if utils.String(r.Data.Postfix) != "" && !strings.HasSuffix(p, utils.String(r.Data.Postfix)) {
			break
		}
		var u *url.URL
		u, err = url.Parse(utils.String(r.Data.Upstream))
		if err != nil {
			break
		}
		u.Path = path.Join(u.Path, utils.String(r.Data.ReplacePrefixWith), p)
		var resp *http.Response
		resp, err = r.Data.clt.Get(u.String())
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
		if utils.Bool(r.Data.CheckMD5) {
			if cmd5 := resp.Header.Get("content-md5"); cmd5 != "" {
				h := md5.New()
				h.Write(b)
				lmd5 := base64.StdEncoding.EncodeToString(h.Sum(nil))
				if lmd5 != cmd5 {
					r.Ctx.StatusCode(http.StatusNotFound)
					err = os.ErrNotExist
					break
				}
			}
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
		if r.Ctx.ClientSupportsGzip() && len(b) > 1024 {
			_, err = r.Ctx.WriteGzip(b)
		} else {
			_, err = r.Ctx.Write(b)
		}
		return err
	}
	return
}

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}
