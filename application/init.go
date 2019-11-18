package application

import (
	"ssh-api/common"
)

type application struct {
	client common.Client
}

func Init(client *common.Client) *application {
	return &application{
		*client,
	}
}
