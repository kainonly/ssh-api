package application

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"ssh-api/common"
)

type listsBody struct {
	Identity []string `json:"identity" validate:"required"`
}

func (app *application) ListsRoute(ctx iris.Context) {
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
	var response []common.GetResponseContent
	for _, identity := range body.Identity {
		content, err := app.client.Get(identity)
		if err != nil {
			ctx.JSON(iris.Map{
				"error": 1,
				"msg":   err.Error(),
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
