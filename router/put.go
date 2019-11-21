package router

import (
	"encoding/base64"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"ssh-api/common"
)

type putBody struct {
	Identity   string `json:"identity" validate:"required"`
	Host       string `json:"host" validate:"required,hostname|ip"`
	Port       uint   `json:"port" validate:"required,numeric"`
	Username   string `json:"username" validate:"required,alphanum"`
	Password   string `json:"password" validate:"required_without=PrivateKey"`
	PrivateKey string `json:"private_key" validate:"required_without=Password,omitempty,base64"`
	Passphrase string `json:"passphrase"`
}

func (r *router) PutRoute(ctx iris.Context) {
	var body putBody
	ctx.ReadJSON(&body)
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
		return
	}
	privateKey, err := base64.StdEncoding.DecodeString(body.PrivateKey)
	if err != nil {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
		return
	}
	if err := r.client.Put(body.Identity, common.ConnectOption{
		Host:       body.Host,
		Port:       body.Port,
		Username:   body.Username,
		Password:   body.Password,
		Key:        privateKey,
		PassPhrase: []byte(body.Passphrase),
	}); err != nil {
		ctx.JSON(iris.Map{
			"error": 1,
			"msg":   err.Error(),
		})
		return
	}
	ctx.JSON(iris.Map{
		"error": 0,
		"msg":   "ok",
	})
}
