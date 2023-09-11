package model

import "github.com/bwmarrin/snowflake"

type Rule struct {
	ID                snowflake.ID `gorm:"primaryKey" json:"id"`
	Prefix            string       `json:"prefix"`
	Postfix           string       `json:"postfix"`
	ReplacePrefixWith string       `json:"replacePrefixWith"`
	Upstream          string       `json:"upstream"`
	CheckMD5          bool         `json:"checkMD5"`
	Active            bool         `json:"active"`
}
