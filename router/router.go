package router

import "ssh-api/client"

type router struct {
	client client.Client
}

func Init(client *client.Client) *router {
	return &router{
		*client,
	}
}
