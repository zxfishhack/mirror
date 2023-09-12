package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"os"
	"strconv"
)

var idGen *snowflake.Node

func init() {
	var nodeId int64 = 0
	var e error
	v := os.Getenv("MIRROR_NODE_ID")
	if v != "" {
		nodeId, e = strconv.ParseInt(v, 10, 63)
		if e != nil {
			log.Fatalf("MIRROR_NODE_ID [%s] parse failed. %v", v, e)
		}
	}
	idGen, e = snowflake.NewNode(nodeId)
	if e != nil {
		log.Fatalf("create id generate with MIRROR_NODE_ID [%s][%d] failed. %v", v, nodeId, e)
	}
}

func ID() snowflake.ID {
	return idGen.Generate()
}

func String(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func StringP(v string) *string {
	return &v
}

func Bool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}
