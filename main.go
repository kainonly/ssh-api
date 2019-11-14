package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"ssh-api/router"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	routes := new(router.Router)
	app.Post("testing", routes.TestingRoute)

	app.Run(iris.Addr("127.0.0.1:3000"), iris.WithoutServerError(iris.ErrServerClosed))
}
