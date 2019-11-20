package client

import (
	"errors"
	"io"
	"net"
	"ssh-api/common"
	"sync"
	"time"
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
	for _, listener := range c.localListener.Map {
		listener.Close()
	}
	if c.localListener != nil {
		c.localListener = newSafeMapListener()
	}
	for _, tunnel := range options {
		go c.setTunnel(identity, tunnel)
	}
	return
}

func (c *Client) setTunnel(identity string, option TunnelOption) {
	addr := common.GetAddr(option.DstIp, option.DstPort)
	remoteAddr := common.GetAddr(option.SrcIp, option.SrcPort)
	localAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		c.error <- common.SendError(identity, err)
		return
	}
	localListener, err := net.ListenTCP("tcp", localAddr);
	if err != nil {
		c.error <- common.SendError(identity, err)
		return
	} else {
		c.localListener.Set(identity, localListener)
	}
	for {
		localConn, err := c.localListener.Get(identity).AcceptTCP()
		if err != nil {
			c.error <- common.SendError(identity, err)
			return
		}
		go c.transport(identity, localConn, remoteAddr)
	}
}

//  tunnel data to the remote server
func (c *Client) transport(identity string, localConn *net.TCPConn, remoteAddr string) {
	defer localConn.Close()
	localConn.SetNoDelay(true)
	localConn.SetKeepAlive(true)
	localConn.SetKeepAlivePeriod(5 * time.Second)
	remoteConn, err := c.runtime[identity].Dial("tcp", remoteAddr)
	if err != nil {
		c.error <- common.SendError(identity, err)
		return
	}
	defer remoteConn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		_, err := io.Copy(localConn, remoteConn)
		if err != nil {
			c.error <- common.SendError(identity, err)
			return
		}
		wg.Done()
	}()
	go func() {
		_, err := io.Copy(remoteConn, localConn)
		if err != nil {
			c.error <- common.SendError(identity, err)
			return
		}
		wg.Done()
	}()
	wg.Wait()
}
