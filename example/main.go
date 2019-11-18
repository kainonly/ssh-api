package main

import (
	"log"
	"net"
	"ssh-api/common"
)

func main() {
	localListener = make(map[string]*net.Listener)
	localConn = make(map[string]*net.Conn)
	//var remoteConn map[string]*net.Conn
	//remoteConn = make(map[string]*net.Conn)
	client, err := connected()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	for _, tunnel := range tunnels {
		localAddr := common.GetAddr(tunnel.DstIp, tunnel.DstPort)
		if listener, err := net.Listen("tcp", localAddr); err == nil {
			localListener[localAddr] = &listener
		} else {
			log.Fatalln(err)
		}
		go func() {
			for {
				if local, err := (*localListener[localAddr]).Accept(); err == nil {
					localConn[localAddr] = &local
				} else {
					log.Fatal(err)
				}

				//go func() {
				//	remoteAddr := common.GetAddr(tunnel.SrcIp, tunnel.SrcPort)
				//	if remote, err := client.Dial("tcp", remoteAddr); err == nil {
				//		remoteConn[remoteAddr] = &remote
				//	} else {
				//		log.Fatal(err)
				//	}
				//	go func() {
				//		io.Copy(*localConn[localAddr], *remoteConn[remoteAddr])
				//	}()
				//	go func() {
				//		io.Copy(*remoteConn[remoteAddr], *localConn[localAddr])
				//	}()
				//}()

			}
		}()
	}

	select {}
}
