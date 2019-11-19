package main

import (
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"ssh-api/common"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	localListener = make(map[string]*net.Listener)
	localConn = make(map[string]*net.Conn)
	remoteConn = make(map[string]*net.Conn)
	client, err := connected()
	if err != nil {
		log.Fatalln(err)
	}
	for _, tunnel := range tunnels {
		go func(tunnel common.TunnelOption) {
			localAddr := common.GetAddr(tunnel.DstIp, tunnel.DstPort)
			if listener, err := net.Listen("tcp", localAddr); err == nil {
				localListener[localAddr] = &listener
			} else {
				log.Fatalln(err)
			}
			for {
				if local, err := (*localListener[localAddr]).Accept(); err == nil {
					localConn[localAddr] = &local
				} else {
					log.Fatal(err)
				}
				remoteAddr := common.GetAddr(tunnel.SrcIp, tunnel.SrcPort)
				if remote, err := client.Dial("tcp", remoteAddr); err == nil {
					remoteConn[remoteAddr] = &remote
				} else {
					log.Fatal(err)
				}
				go io.Copy(*localConn[localAddr], *remoteConn[remoteAddr])
				go io.Copy(*remoteConn[remoteAddr], *localConn[localAddr])
			}
		}(tunnel)
	}
	select {}
}
