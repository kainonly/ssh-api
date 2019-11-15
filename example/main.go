package main

import (
	"log"
	"ssh-api/common"
)

func main() {
	option, err := GetOption("./debug.json")
	if err != nil {
		log.Fatalln(err)
	}
	c := new(container.Client)
	client, err := c.Testing(option)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}
