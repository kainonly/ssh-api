package client

import (
	"golang.org/x/crypto/ssh"
)

type Client struct {
	options       map[string]*ConnectOption
	runtime       map[string]*ssh.Client
	localListener *safeMapListener
}

// Inject ssh client service
func InjectClient() *Client {
	return &Client{
		options:       make(map[string]*ConnectOption),
		runtime:       make(map[string]*ssh.Client),
		localListener: newSafeMapListener(),
	}
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*ConnectOption {
	return c.options
}
