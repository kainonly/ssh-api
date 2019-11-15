package main

import (
	"ssh-api/common"
	"ssh-api/dev"
)

func main() {
	option, err := dev.TestOption()
	if err != nil {
		println(err.Error())
		return
	}
	c := new(common.Client)
	client, err := c.Testing(option)
	if err != nil {
		println(err.Error())
		return
	}
	defer client.Close()
	println(string(client.ClientVersion()))
}
