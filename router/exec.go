package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type execBody struct {
	Identity string `json:"identity" validate:"required"`
	Bash     string `json:"bash" validate:"required"`
}

func (r *router) ExecRoute(ctx iris.Context) {
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
	output, err := r.client.Exec(body.Identity, body.Bash)
	if err == nil {
		ctx.JSON(iris.Map{
			"error": 0,
			"data":  string(output),
		})
	} else {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
	}
}
