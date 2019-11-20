package client

import (
	"golang.org/x/crypto/ssh"
	"ssh-api/common"
)

type Client struct {
	options       map[string]*ConnectOption
	runtime       map[string]*ssh.Client
	localListener *safeMapListener
	error         chan common.Error
}

// Inject ssh client service
func InjectClient() *Client {
	channel := make(chan common.Error)
	go lisenErrorChannel(channel)
	return &Client{
		options:       make(map[string]*ConnectOption),
		runtime:       make(map[string]*ssh.Client),
		localListener: newSafeMapListener(),
		error:         channel,
	}
}

func lisenErrorChannel(channel chan common.Error) {
	defer func() {
		if r := recover(); r != nil {
			println(r.(string))
		}
	}()
	select {
	case data := <-channel:
		panic("<" + data.Name + "> " + data.Error())
	}
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*ConnectOption {
	return c.options
}
