package client

import (
	"golang.org/x/crypto/ssh"
	"net"
)

type Client struct {
	runtime map[string]*ssh.Client
	server  map[string]*net.TCPListener
	options map[string]*ConnectOption
	Error   chan error
}

// Inject ssh client service
func InjectClient() *Client {
	return &Client{
		runtime: make(map[string]*ssh.Client),
		server:  make(map[string]*net.TCPListener),
		options: make(map[string]*ConnectOption),
		Error:   make(chan error),
	}
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*ConnectOption {
	return c.options
}
