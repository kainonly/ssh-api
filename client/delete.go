package client

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
