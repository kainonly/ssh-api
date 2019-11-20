package client

import (
	"errors"
	"io"
	"net"
	"ssh-api/common"
	"sync"
)

type TunnelOption struct {
	SrcIp   string `json:"src_ip" validate:"required,ip"`
	SrcPort uint   `json:"src_port" validate:"required,numeric"`
	DstIp   string `json:"dst_ip" validate:"required,ip"`
	DstPort uint   `json:"dst_port" validate:"required,numeric"`
}

func (c *Client) Tunnels(identity string, options []TunnelOption) (err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	for _, conn := range c.localConn.Map {
		(*conn).Close()
	}
	c.localConn = newSafeMapConn()
	for _, listener := range c.localListener.Map {
		(*listener).Close()
	}
	c.localListener = newSafeMapListener()
	for _, tunnel := range options {
		go c.mutilTunnel(identity, tunnel)
	}
	return
}

func (c *Client) mutilTunnel(identity string, option TunnelOption) {
	localAddr := common.GetAddr(option.DstIp, option.DstPort)
	remoteAddr := common.GetAddr(option.SrcIp, option.SrcPort)
	localListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		println("<" + identity + ">:" + err.Error())
		return
	} else {
		c.localListener.Set(identity, &localListener)
	}
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			println("<" + identity + ">:" + err.Error())
			return
		} else {
			c.localConn.Set(identity, &localConn)
		}
		go c.forward(identity, remoteAddr)
	}
}

//  tunnel data to the remote server
func (c *Client) forward(identity string, remoteAddr string) {
	localConn := *c.localConn.Get(identity)
	defer localConn.Close()
	remoteConn, err := c.runtime[identity].Dial("tcp", remoteAddr)
	if err != nil {
		println("remote <" + identity + ">:" + err.Error())
		return
	}
	defer remoteConn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func(local net.Conn, remote net.Conn) {
		io.Copy(remote, local)
		wg.Done()
	}(localConn, remoteConn)
	go func(local net.Conn, remote net.Conn) {
		io.Copy(local, remote)
		wg.Done()
	}(localConn, remoteConn)
	wg.Wait()
}
