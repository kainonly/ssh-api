package client

import (
	"golang.org/x/crypto/ssh"
	"ssh-api/common"
)

type ConnectOption struct {
	Host       string
	Port       uint
	Username   string
	Password   string
	Key        []byte
	PassPhrase []byte
}

// Generate AuthMethod
func (c *Client) authMethod(option ConnectOption) (auth []ssh.AuthMethod, err error) {
	if option.Key == nil {
		// Password AuthMethod
		auth = []ssh.AuthMethod{
			ssh.Password(option.Password),
		}
	} else {
		// PrivateKey AuthMethod
		var signer ssh.Signer
		if option.PassPhrase != nil {
			// With Passphrase
			if signer, err = ssh.ParsePrivateKeyWithPassphrase(
				option.Key,
				option.PassPhrase,
			); err != nil {
				return
			}
		} else {
			// Without Passphrase
			if signer, err = ssh.ParsePrivateKey(
				option.Key,
			); err != nil {
				return
			}
		}
		auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	}
	return
}

// Ssh client connection
func (c *Client) connect(option ConnectOption) (client *ssh.Client, err error) {
	auth, err := c.authMethod(option)
	if err != nil {
		return
	}
	config := ssh.ClientConfig{
		User:            option.Username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := common.GetAddr(option.Host, option.Port)
	client, err = ssh.Dial("tcp", addr, &config)
	return
}
