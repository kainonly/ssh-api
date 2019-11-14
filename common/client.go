package common

import (
	"golang.org/x/crypto/ssh"
)

type Client struct{}

type TestingOption struct {
	Host     string
	Port     uint
	Username string
	Key      []byte
}

func (_ *Client) Testing(option TestingOption) (client *ssh.Client, err error) {
	signer, err := ssh.ParsePrivateKey(option.Key)
	config := &ssh.ClientConfig{
		User: option.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err = ssh.Dial("tcp", option.Host, config)
	return
}
