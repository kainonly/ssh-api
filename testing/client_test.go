package main

import (
	"ssh-api/common"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	option, err := GetOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := common.InjectClient()
	client, err := c.Testing(option)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	println(string(client.ClientVersion()))
}

func TestPut(t *testing.T) {
	option, err := GetOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := common.InjectClient()
	err = c.Put("test", option)
	if err != nil {
		t.Error(err)
	}
	content, err := c.Get("test")
	if err != nil {
		t.Error(err)
	}
	println(content.Identity)
	println(content.Connected)
}

func TestExec(t *testing.T) {
	option, err := GetOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := common.InjectClient()
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
	option, err := GetOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := common.InjectClient()
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
