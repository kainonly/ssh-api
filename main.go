package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"net/http"
	_ "net/http/pprof"
	"ssh-api/client"
	"ssh-api/common"
	"ssh-api/router"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	common.InitLevelDB("data")
	common.InitBufPool()
	routes := router.Init(
		client.InjectClient(),
	)
	app.Post("testing", routes.TestingRoute)
	app.Post("put", routes.PutRoute)
	app.Post("exec", routes.ExecRoute)
	app.Post("delete", routes.DeleteRoute)
	app.Post("get", routes.GetRoute)
	app.Post("all", routes.AllRoute)
	app.Post("lists", routes.ListsRoute)
	app.Post("tunnels", routes.TunnelsRoute)
	app.Run(iris.Addr(":3000"), iris.WithoutServerError(iris.ErrServerClosed))
}
