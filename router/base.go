package router

import "ssh-api/service"

type application struct {
	client service.Client
}

func Container(client *service.Client) *application {
	return &application{
		*client,
	}
}
