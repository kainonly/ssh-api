package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type deleteBody struct {
	Identity string `json:"identity" validate:"required"`
}

func (app *application) DeleteRoute(ctx iris.Context) {
	var body deleteBody
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
