package application

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"ssh-api/common"
)

type tunnelsBody struct {
	Identity string                `json:"identity" validate:"required"`
	Tunnels  []common.TunnelOption `json:"tunnels" validate:"required"`
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
	err := app.client.Tunnels(body.Identity, body.Tunnels)
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
