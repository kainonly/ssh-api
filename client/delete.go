package client

// Delete ssh client
func (c *Client) Delete(identity string) (err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		return
	}
	c.closeTunnel(identity)
	if c.runtime[identity] != nil {
		c.runtime[identity].Close()
	}
	delete(c.runtime, identity)
	delete(c.options, identity)
	return
}
