package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type execBody struct {
	Identity string `json:"identity" validate:"required"`
	Bash     string `json:"bash" validate:"required"`
}

func (app *application) ExecRoute(ctx iris.Context) {
	var body execBody
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
