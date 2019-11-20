package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"ssh-api/client"
)

type listsBody struct {
	Identity []string `json:"identity" validate:"required"`
}

func (r *router) ListsRoute(ctx iris.Context) {
	var body listsBody
	ctx.ReadJSON(&body)
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
		return
	}
	var response []client.Information
	for _, identity := range body.Identity {
		content, err := r.client.Get(identity)
		if err != nil {
			ctx.JSON(iris.Map{
				"error": 1,
				"msg":   identity + ":" + err.Error(),
			})
			return
		}
		response = append(response, content)
	}
	ctx.JSON(iris.Map{
		"error": 0,
		"data":  response,
	})
}
