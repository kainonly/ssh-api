package client

import (
	"golang.org/x/crypto/ssh"
)

type Client struct {
	runtime       map[string]*ssh.Client
	options       map[string]*ConnectOption
	tunnels       map[string]*[]TunnelOption
	localListener *safeMapListener
	localConn     *safeMapConn
	remoteConn    *safeMapConn
}

// Inject ssh client service
func InjectClient() *Client {
	return &Client{
		runtime:       make(map[string]*ssh.Client),
		options:       make(map[string]*ConnectOption),
		tunnels:       make(map[string]*[]TunnelOption),
		localListener: newSafeMapListener(),
		localConn:     newSafeMapConn(),
		remoteConn:    newSafeMapConn(),
	}
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*ConnectOption {
	return c.options
}
