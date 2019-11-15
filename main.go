package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"ssh-api/router"
	"ssh-api/service"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	route := router.Container(
		service.InjectClient(),
	)
	app.Post("testing", route.TestingRoute)
	app.Post("put", route.PutRoute)
	app.Post("exec", route.ExecRoute)
	app.Post("delete", route.DeleteRoute)
	app.Post("get", route.GetRoute)
	app.Post("all", route.AllRoute)
	app.Post("lists", route.ListsRoute)
	app.Post("tunnels", route.TunnelsRoute)
	app.Run(iris.Addr("127.0.0.1:3000"), iris.WithoutServerError(iris.ErrServerClosed))
}
