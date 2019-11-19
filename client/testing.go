package client

import "golang.org/x/crypto/ssh"

// Test ssh client connection
func (c *Client) Testing(option ConnectOption) (sshClient *ssh.Client, err error) {
	return c.connect(option)
}
