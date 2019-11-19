package client

import "sync"

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
