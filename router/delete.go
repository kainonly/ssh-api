package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type deleteBody struct {
	Identity string `json:"identity" validate:"required"`
}

func (r *router) DeleteRoute(ctx iris.Context) {
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
	err := r.client.Delete(body.Identity)
	if err == nil {
		ctx.JSON(iris.Map{
			"error": 0,
			"msg":   "ok",
		})
	} else {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
	}
}
