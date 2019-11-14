package common

type Common struct {
	Config
	Client
}

func New() *Common {
	return &Common{}
}
