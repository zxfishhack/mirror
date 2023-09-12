package rule

import (
	"context"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zxfishhack/mirror/pkg/model"
	"github.com/zxfishhack/mirror/pkg/storage"
	"github.com/zxfishhack/mirror/pkg/utils"
	"gorm.io/gorm"
	"log"
	"sync"
)

type manager struct {
	db     *gorm.DB
	addr   string
	mtx    sync.Mutex
	app    *iris.Application
	create storage.CreateStorageFunc
}

type ManagerController struct {
	Ctx iris.Context

	Data *manager
}

func ManagerConfigure(addr string, db *gorm.DB, create storage.CreateStorageFunc) func(app *mvc.Application) {
	return func(app *mvc.Application) {
		m := &ManagerController{Data: &manager{
			addr:   addr,
			db:     db,
			create: create,
		},
		}
		app.Handle(m)

		err := m.reconfigure()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (m *ManagerController) reconfigure() (err error) {
	m.Data.mtx.Lock()
	defer m.Data.mtx.Unlock()
	rules, err := m.Data.db.Model(&model.Rule{}).Rows()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		return
	}
	if m.Data.app != nil {
		m.Data.app.Shutdown(context.Background())
	}
	m.Data.app = iris.New()
	active := make([]snowflake.ID, 0)
	inactive := make([]snowflake.ID, 0)
	for rules.Next() {
		var r model.Rule
		if err = m.Data.db.ScanRows(rules, &r); err == nil {
			var c *RuleController
			c, err = newRule(r, m.Data.create)
			if err != nil {
				log.Printf("apply rule[%#v] failed %v", r, err)
				err = nil
				inactive = append(inactive, r.ID)
				continue
			}
			mvc.Configure(m.Data.app.Party(utils.String(r.Prefix)), func(app *mvc.Application) {
				app.Handle(c)
			})
			active = append(active, r.ID)
		} else {
			break
		}
	}
	m.Data.db.Model(&model.Rule{}).Where("id IN ?", active).UpdateColumn("active", true)
	m.Data.db.Model(&model.Rule{}).Where("id IN ?", inactive).UpdateColumn("active", false)
	go m.Data.app.Run(iris.Addr(m.Data.addr),
		iris.WithRemoteAddrHeader("X-Real-Ip"),
		iris.WithoutRemoteAddrHeader("X-Forwarded-For"),
	)
	return
}

func (m *ManagerController) GetReconfigure() error {
	return m.reconfigure()
}

func (m *ManagerController) GetRules() (res []model.Rule, err error) {
	rules, err := m.Data.db.Model(&model.Rule{}).Rows()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		return
	}
	for rules.Next() {
		var r model.Rule
		if err = m.Data.db.ScanRows(rules, &r); err == nil {
			res = append(res, r)
		}
	}
	return
}

func (m *ManagerController) PostRule() (err error) {
	var r model.Rule
	err = m.Ctx.ReadJSON(&r)
	if err != nil {
		return
	}
	r.ID = utils.ID()
	r.Normalize()
	return m.Data.db.Model(&model.Rule{}).Create(&r).Error
}

func (m *ManagerController) PostRuleBy(id int64) (err error) {
	var r model.Rule
	err = m.Ctx.ReadJSON(&r)
	if err != nil {
		return
	}
	r.ID = snowflake.ID(id)
	r.Normalize()
	return m.Data.db.Model(&model.Rule{}).Where(&model.Rule{ID: r.ID}).Updates(&r).Error
}

func (m *ManagerController) DeleteRuleBy(id int64) (err error) {
	return m.Data.db.Model(&model.Rule{}).Delete(&model.Rule{ID: snowflake.ID(id)}).Error
}
