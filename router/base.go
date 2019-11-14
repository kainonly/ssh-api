package router

import "ssh-api/common"

type Router struct {
	common.Config
	common.Client
}

func New(common *common.Common) *Router {
	return &Router{
		common.Config,
		common.Client,
	}
}
