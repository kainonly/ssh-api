package application

import "ssh-api/client"

type application struct {
	client client.Client
}

func Init(client *client.Client) *application {
	return &application{
		*client,
	}
}
