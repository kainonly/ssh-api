package main

import (
	"ssh-api/service"
	"testing"
)

func TestSimple(t *testing.T) {
	option, err := GetOption("../debug.json")
	if err != nil {
		t.Error(err)
	}
	c := new(service.Client)
	client, err := c.Testing(option)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	println(string(client.ClientVersion()))
}

func TestPut(t *testing.T) {

}