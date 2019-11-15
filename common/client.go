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

// Factory SSH AuthMethod
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

// SSH Connect Testing
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
	addr := option.Host + ":" + strconv.Itoa(int(option.Port))
	client, err = ssh.Dial("tcp", addr, &config)
	return
}

func (c *Client) Put(identity string, option ConnectOption) {

}

func (c *Client) Get(identity string) (exists bool, result interface{}) {
	return true, nil
}
