package client

import (
	"errors"
	"io"
	"net"
	"strconv"
	"sync"
)

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
