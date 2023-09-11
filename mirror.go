package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/zxfishhack/mirror/pkg/console"
	"github.com/zxfishhack/mirror/pkg/model"
	"github.com/zxfishhack/mirror/pkg/rule"
	"github.com/zxfishhack/mirror/pkg/storage"
	"github.com/zxfishhack/mirror/pkg/storage/simple"
)

//go:generate go-bindata -pkg console -prefix "pkg/console/dist" -o pkg/console/assets.go pkg/console/dist/...

func main() {
	var storagePath, internalPath, storageType, addr string
	var consoleAddr string
	flag.StringVar(&storagePath, "f", "/files", "mirror files dir")
	flag.StringVar(&internalPath, "s", "/data", "internal storage dir")
	flag.StringVar(&storageType, "type", "simple", "storage type: simple")
	flag.StringVar(&addr, "l", ":80", "listen address")
	flag.StringVar(&consoleAddr, "console", ":8080", "console listen address")
	flag.Parse()

	var createFunc storage.CreateStorageFunc

	switch storageType {
	case "simple":
		createFunc = simple.CreateFunc(storagePath)
	default:
		flag.PrintDefaults()
		fmt.Printf("%s storage type not support.", storageType)
		return
	}

	db, err := model.New(internalPath)
	if err != nil {
		fmt.Printf("init internal data failed %v", err)
		return
	}

	app := iris.New()
	mvc.Configure(app.Party("/"), console.Handle)
	mvc.Configure(app.Party("/"), rule.ManagerConfigure(addr, db, createFunc))

	app.Use(recover.New())

	var options []iris.Configurator

	options = append(options, iris.WithoutServerError(iris.ErrServerClosed))

	app.Run(iris.Addr(consoleAddr), options...)
}
