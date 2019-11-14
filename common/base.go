package common

type Common struct {
	Config
	Client
}

func New() *Common {
	common := Common{}
	common.Config.init()
	return &common
}
