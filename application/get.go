package application

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type getBody struct {
	Identity string `json:"identity" validate:"required"`
}

func (app *application) GetRoute(ctx iris.Context) {
	var body getBody
	ctx.ReadJSON(&body)
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
		return
	}
	content, err := app.client.Get(body.Identity)
	if err == nil {
		ctx.JSON(iris.Map{
			"error": 0,
			"data":  content,
		})
	} else {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
	}
}
