package client

import (
	"golang.org/x/crypto/ssh"
)

type Client struct {
	runtime map[string]*ssh.Client
	options map[string]*ConnectOption
}

type (
	ConnectOptionWithIdentity struct {
		Identity string
		ConnectOption
	}
	TunnelOption struct {
		SrcIp   string `json:"src_ip" validate:"required"`
		SrcPort uint   `json:"src_port" validate:"required"`
		DstIp   string `json:"dst_ip" validate:"required"`
		DstPort uint   `json:"dst_port" validate:"required"`
	}
)

// Inject ssh client service
func InjectClient() *Client {
	client := Client{}
	client.runtime = make(map[string]*ssh.Client)
	client.options = make(map[string]*ConnectOption)
	return &client
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*ConnectOption {
	return c.options
}
