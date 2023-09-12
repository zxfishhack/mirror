package model

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zxfishhack/mirror/pkg/utils"
	"strings"
)

type Rule struct {
	ID                snowflake.ID `gorm:"primaryKey" json:"id"`
	Prefix            *string      `gorm:"unique" json:"prefix"`
	Postfix           *string      `json:"postfix"`
	ProxyUrl          *string      `json:"proxyUrl"`
	ReplacePrefixWith *string      `json:"replacePrefixWith"`
	Upstream          *string      `json:"upstream"`
	CheckMD5          *bool        `json:"checkMD5"`
	Active            *bool        `json:"active"`
}

func (r *Rule) Normalize() {
	r.Prefix = utils.StringP(strings.Trim(utils.String(r.Prefix), "/") + "/")
	r.ReplacePrefixWith = utils.StringP(strings.Trim(utils.String(r.ReplacePrefixWith), "/"))
	r.Upstream = utils.StringP(strings.Trim(utils.String(r.Upstream), "/") + "/")
}
