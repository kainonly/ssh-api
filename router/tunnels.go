package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type tunnelsBody struct {
	Identity string     `json:"identity" validate:"required"`
	Tunnels  [][]string `json:"tunnels" validate:"required"`
}

func (app *application) TunnelsRoute(ctx iris.Context) {
	var body tunnelsBody
	ctx.ReadJSON(&body)
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
		return
	}
	//
	ctx.JSON(iris.Map{
		"error": 0,
		"msg":   "ok",
	})
}
