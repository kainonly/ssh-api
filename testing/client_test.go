package testing

import (
	"ssh-api/client"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	option, err := GetDebugOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := client.InjectClient()
	sshclient, err := c.Testing(option)
	if err != nil {
		t.Error(err)
	}
	defer sshclient.Close()
	println(string(sshclient.ClientVersion()))
}

func TestPut(t *testing.T) {
	option, err := GetDebugOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := client.InjectClient()
	err = c.Put("test", option)
	if err != nil {
		t.Error(err)
	}
	content, err := c.Get("test")
	if err != nil {
		t.Error(err)
	}
	println(content.Identity)
	println(content.Host)
}

func TestExec(t *testing.T) {
	option, err := GetDebugOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := client.InjectClient()
	err = c.Put("test", option)
	if err != nil {
		t.Error(err)
	}
	output, err := c.Exec("test", "uptime")
	if err != nil {
		t.Error(err)
	}
	println(string(output))
}

func TestMultiExec(t *testing.T) {
	option, err := GetDebugOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := client.InjectClient()
	err = c.Put("test", option)
	if err != nil {
		t.Error(err)
	}
	output, err := c.Exec("test", "uptime")
	if err != nil {
		t.Error(err)
	}
	println(string(output))
	time.Sleep(3 * time.Second)
	output, err = c.Exec("test", "ls -l")
	if err != nil {
		t.Error(err)
	}
	println(string(output))
}

func TestDelete(t *testing.T) {
	option, err := GetDebugOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := client.InjectClient()
	err = c.Put("test", option)
	if err != nil {
		t.Error(err)
	}
	err = c.Delete("test")
	if err != nil {
		t.Error(err)
	}
}

func TestAll(t *testing.T) {
	option, err := GetDebugOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := client.InjectClient()
	err = c.Put("test", option)
	if err != nil {
		t.Error(err)
	}
	err = c.Put("abc", option)
	if err != nil {
		t.Error(err)
	}
	var keys []string
	for key, _ := range c.GetClientOptions() {
		keys = append(keys, key)
	}
	println(keys)
}

func TestLists(t *testing.T) {
	option, err := GetDebugOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := client.InjectClient()
	err = c.Put("test", option)
	if err != nil {
		t.Error(err)
	}
	err = c.Put("abc", option)
	if err != nil {
		t.Error(err)
	}
	var response []client.Information
	for _, identity := range []string{"test", "abc"} {
		content, err := c.Get(identity)
		if err != nil {
			t.Error(err)
		}
		response = append(response, content)
	}
}
