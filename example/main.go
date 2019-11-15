package main

import (
	"io"
	"log"
	"net"
	"ssh-api/service"
)

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
	listener, err := net.Listen("tcp", "127.0.0.1:5601")
	if err != nil {
		log.Fatal(err)
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
			println(remote)
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
