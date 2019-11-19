package common

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
)

type (
	Client struct {
		runtime map[string]*ssh.Client
		options map[string]*ConnectOption
	}
	ConnectOption struct {
		Host       string
		Port       uint
		Username   string
		Password   string
		Key        []byte
		PassPhrase []byte
	}
	ConnectOptionWithIdentity struct {
		Identity string
		ConnectOption
	}
	Information struct {
		Identity  string         `json:"identity"`
		Host      string         `json:"host"`
		Port      uint           `json:"port"`
		Username  string         `json:"username"`
		Connected string         `json:"connected"`
		Tunnels   []TunnelOption `json:"tunnels"`
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

// Get Addr
func GetAddr(ip string, port uint) string {
	return ip + ":" + strconv.Itoa(int(port))
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
		Timeout:         time.Duration(3 * time.Second),
	}
	addr := GetAddr(option.Host, option.Port)
	client, err = ssh.Dial("tcp", addr, &config)
	return
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*ConnectOption {
	return c.options
}

// Test ssh client connection
func (c *Client) Testing(option ConnectOption) (client *ssh.Client, err error) {
	return c.connect(option)
}

// Add or modify the ssh client
func (c *Client) Put(identity string, option ConnectOption) (err error) {
	err = c.Delete(identity)
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.options[identity] = &option
		c.runtime[identity], err = c.connect(option)
	}()
	wg.Wait()
	return
}

// Get ssh client information
func (c *Client) Get(identity string) (content Information, err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	option := c.options[identity]
	content = Information{
		Identity:  identity,
		Host:      option.Host,
		Port:      option.Port,
		Username:  option.Username,
		Connected: string(c.runtime[identity].ClientVersion()),
	}
	return
}

// Remotely execute commands via SSH
func (c *Client) Exec(identity string, cmd string) (output []byte, err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		session, err := c.runtime[identity].NewSession()
		if err != nil {
			return
		}
		defer session.Close()
		output, err = session.Output(cmd)
	}()
	wg.Wait()
	return
}

// Delete ssh client
func (c *Client) Delete(identity string) (err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		return
	}
	err = c.runtime[identity].Close()
	if err != nil {
		return
	}
	delete(c.runtime, identity)
	delete(c.options, identity)
	return
}

func (c *Client) Tunnels(identity string, options []TunnelOption) (err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(options))
	for _, option := range options {
		go func() {
			dstAddr := option.DstIp + ":" + strconv.Itoa(int(option.DstPort))
			listener, err := net.Listen("tcp", dstAddr)
			if err != nil {
				return
			}
			defer listener.Close()
			wg.Done()
			go func() {
				for {
					local, err := listener.Accept()
					if err != nil {
						return
					}
					go func() {
						srcAddr := option.SrcIp + ":" + strconv.Itoa(int(option.SrcPort))
						remote, err := c.runtime[identity].Dial("tcp", srcAddr)
						if err != nil {
							return
						}
						go func() {
							_, err := io.Copy(local, remote)
							if err != nil {
								return
							}
						}()
						go func() {
							_, err := io.Copy(remote, local)
							if err != nil {
								return
							}
						}()
					}()
				}
			}()
		}()
	}
	wg.Wait()
	return
}
