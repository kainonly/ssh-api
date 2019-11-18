package application

import (
	"github.com/kataras/iris/v12"
)

func (app *application) AllRoute(ctx iris.Context) {
	var keys []string
	for key := range app.client.GetClientOptions() {
		keys = append(keys, key)
	}
	ctx.JSON(iris.Map{
		"error": 0,
		"data":  keys,
	})
}
