package router

import (
	"github.com/kataras/iris/v12"
)

func (m *Router) Put(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
}
