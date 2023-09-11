package model

import (
	"github.com/bwmarrin/snowflake"
	"strings"
)

type Rule struct {
	ID                snowflake.ID `gorm:"primaryKey" json:"id"`
	Prefix            string       `gorm:"unique" json:"prefix"`
	Postfix           string       `json:"postfix"`
	ReplacePrefixWith string       `json:"replacePrefixWith"`
	Upstream          string       `json:"upstream"`
	CheckMD5          bool         `json:"checkMD5"`
	Active            bool         `json:"active"`
}

func (r *Rule) Normalize() {
	r.Prefix = strings.Trim(r.Prefix, "/") + "/"
	r.ReplacePrefixWith = strings.Trim(r.ReplacePrefixWith, "/") + "/"
	r.Upstream = strings.Trim(r.Upstream, "/") + "/"
}
