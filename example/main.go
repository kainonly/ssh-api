package main

import (
	"io"
	"log"
	"net"
	"ssh-api/service"
)

func NewLocalListener(address string) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", address)
	return
}

func main() {
	option, err := GetOption("./debug.json")
	if err != nil {
		log.Fatalln(err)
	}
	c := new(service.Client)
	client, err := c.Testing(option)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	listener, err := NewLocalListener(":5601")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		local, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			remote, err := client.Dial("tcp", "0.0.0.0:5601")
			if err != nil {
				log.Fatal(err)
			}
			go func() {
				_, err := io.Copy(local, remote)
				if err != nil {
					log.Fatal(err)
				}
			}()
			go func() {
				_, err := io.Copy(remote, local)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}()
	}
}
