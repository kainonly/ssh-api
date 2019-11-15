package common

import (
	"golang.org/x/crypto/ssh"
	"strconv"
)

type Client struct {
	Options map[string]struct {
		Host     string
		Port     uint
		Username string
	}
	Runtime map[string]*ssh.Client
}

type ConnectOption struct {
	Host       string
	Port       uint
	Username   string
	Password   string
	Key        []byte
	PassPhrase []byte
}

func (c *Client) authMethod(option ConnectOption) (auth []ssh.AuthMethod, err error) {
	if option.Key == nil {
		auth = []ssh.AuthMethod{
			ssh.Password(option.Password),
		}
	} else {
		var signer ssh.Signer
		if option.PassPhrase != nil {
			if signer, err = ssh.ParsePrivateKeyWithPassphrase(
				option.Key,
				option.PassPhrase,
			); err != nil {
				return
			}
		} else {
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

func (c *Client) Testing(option ConnectOption) (client *ssh.Client, err error) {
	auth, err := c.authMethod(option)
	if err != nil {
		return
	}
	config := ssh.ClientConfig{
		User:            option.Username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err = ssh.Dial(
		"tcp",
		option.Host+":"+strconv.Itoa(int(option.Port)),
		&config,
	)
	return
}

func (c *Client) Get(identity string) (exists bool, result interface{}) {
	return true, nil
}
