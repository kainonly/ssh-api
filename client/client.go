package client

import (
	"golang.org/x/crypto/ssh"
	"ssh-api/common"
)

type Client struct {
	options       map[string]*common.ConnectOption
	tunnels       map[string]*[]common.TunnelOption
	runtime       map[string]*ssh.Client
	localListener *safeMapListener
	localConn     *safeMapConn
	remoteConn    *safeMapConn
}

// Inject ssh client service
func InjectClient() *Client {
	return &Client{
		options:       make(map[string]*common.ConnectOption),
		tunnels:       make(map[string]*[]common.TunnelOption),
		runtime:       make(map[string]*ssh.Client),
		localListener: newSafeMapListener(),
		localConn:     newSafeMapConn(),
		remoteConn:    newSafeMapConn(),
	}
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*common.ConnectOption {
	return c.options
}
