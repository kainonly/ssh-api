package client

import (
	"sync"
)

type ConnectOptionWithIdentity struct {
	Identity string
	ConnectOption
}

// Add or modify the ssh client
func (c *Client) Put(identity string, option ConnectOption) (err error) {
	err = c.Delete(identity)
	if err != nil {
		return
	}
	c.closeTunnel(identity)
	if c.runtime[identity] != nil {
		c.runtime[identity].Close()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.options[identity] = &option
		c.runtime[identity], err = c.connect(option)
		if err != nil {
			return
		}
	}()
	wg.Wait()
	return
}
