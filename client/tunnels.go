package client

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"ssh-api/common"
	"sync"
	"time"
)

type TunnelOption struct {
	SrcIp   string `json:"src_ip" validate:"required"`
	SrcPort uint   `json:"src_port" validate:"required"`
	DstIp   string `json:"dst_ip" validate:"required"`
	DstPort uint   `json:"dst_port" validate:"required"`
}

func (c *Client) Tunnels(identity string, options []TunnelOption) (err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	for _, tunnel := range options {
		go c.setTunnel(
			c.runtime[identity],
			c.server[identity],
			tunnel,
		)
	}
	return
}

func (c *Client) setTunnel(client *ssh.Client, listener *net.TCPListener, tunnel TunnelOption) {
	addr := common.GetAddr(tunnel.DstIp, tunnel.DstPort)
	remoteAddr := common.GetAddr(tunnel.SrcIp, tunnel.SrcPort)
	localAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		c.Error <- err
		return
	}
	listener, err = net.ListenTCP("tcp", localAddr)
	if err != nil {
		c.Error <- err
		return
	}
	for {
		localConn, err := listener.AcceptTCP()
		if err != nil {
			c.Error <- err
			return
		}
		go c.transport(
			localConn,
			client,
			remoteAddr,
		)
	}
}

//  tunnel data to the remote server
func (c *Client) transport(localConn *net.TCPConn, client *ssh.Client, remoteAddr string) {
	defer localConn.Close()
	localConn.SetNoDelay(true)
	localConn.SetKeepAlive(true)
	localConn.SetKeepAlivePeriod(5 * time.Second)
	remoteConn, err := client.Dial("tcp", remoteAddr)
	if err != nil {
		c.Error <- err
		return
	}
	defer remoteConn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		_, err = io.Copy(localConn, remoteConn)
		if err != nil {
			c.Error <- err
			return
		}
		wg.Done()
	}()
	go func() {
		_, err = io.Copy(remoteConn, localConn)
		if err != nil {
			c.Error <- err
			return
		}
		wg.Done()
	}()
	wg.Wait()
	return
}
