package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"ssh-api/application"
	"ssh-api/client"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	handlers := application.Init(
		client.InjectClient(),
	)
	app.Post("testing", handlers.TestingRoute)
	app.Post("put", handlers.PutRoute)
	app.Post("exec", handlers.ExecRoute)
	app.Post("delete", handlers.DeleteRoute)
	app.Post("get", handlers.GetRoute)
	app.Post("all", handlers.AllRoute)
	app.Post("lists", handlers.ListsRoute)
	app.Post("tunnels", handlers.TunnelsRoute)
	app.Run(iris.Addr(":3000"), iris.WithoutServerError(iris.ErrServerClosed))
}
