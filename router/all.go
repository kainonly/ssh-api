package router

import (
	"github.com/kataras/iris/v12"
)

func (app *Router) AllRoute(ctx iris.Context) {
	//
	ctx.JSON(iris.Map{
		"error": 0,
		"msg":   "ok",
	})
}