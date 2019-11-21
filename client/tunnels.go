package client

import (
	"errors"
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
	c.closeTunnel(identity)
	for _, tunnel := range options {
		go c.mutilTunnel(identity, tunnel)
	}
	return
}

func (c *Client) closeTunnel(identity string) {
	for _, conn := range c.remoteConn.Map[identity] {
		(*conn).Close()
	}
	c.remoteConn.Clear(identity)
	for _, conn := range c.localConn.Map[identity] {
		(*conn).Close()
	}
	c.localConn.Clear(identity)
	for _, listener := range c.localListener.Map[identity] {
		(*listener).Close()
	}
	c.localListener.Clear(identity)
}

func (c *Client) mutilTunnel(identity string, option TunnelOption) {
	localAddr := common.GetAddr(option.DstIp, option.DstPort)
	remoteAddr := common.GetAddr(option.SrcIp, option.SrcPort)
	localListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		println("<" + identity + ">:" + err.Error())
		return
	} else {
		c.localListener.Set(identity, localAddr, &localListener)
	}
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			println("<" + identity + ">:" + err.Error())
			return
		} else {
			c.localConn.Set(identity, localAddr, &localConn)
		}
		go c.forward(identity, localAddr, remoteAddr)
	}
}

//  tunnel data to the remote server
func (c *Client) forward(identity string, localAddr string, remoteAddr string) {
	localConn := *c.localConn.Get(identity, localAddr)
	defer localConn.Close()
	remoteConn, err := c.runtime[identity].Dial("tcp", remoteAddr)
	if err != nil {
		println("remote <" + identity + ">:" + err.Error())
		return
	} else {
		c.remoteConn.Set(identity, localAddr, &remoteConn)
	}
	defer remoteConn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		common.Copy(remoteConn, localConn)
	}()
	go func() {
		defer wg.Done()
		common.Copy(localConn, remoteConn)
	}()
	println("<finish-1>")
	wg.Wait()
	println("<finish-2>")
}
