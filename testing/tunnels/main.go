package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"ssh-api/client"
	"ssh-api/common"
	"ssh-api/testing"
)

var (
	localListener map[string]*net.Listener
	tunnels       = []client.TunnelOption{
		{
			SrcIp:   "127.0.0.1",
			SrcPort: 5601,
			DstIp:   "127.0.0.1",
			DstPort: 5601,
		},
		{
			SrcIp:   "127.0.0.1",
			SrcPort: 9200,
			DstIp:   "127.0.0.1",
			DstPort: 9200,
		},
	}
)

func connected() (sshClient *ssh.Client, err error) {
	option, err := testing.GetDebugOption("./debug.json")
	if err != nil {
		return
	}
	c := client.InjectClient()
	return c.Testing(option)
}

func setTunnel(client *ssh.Client, tunnel client.TunnelOption) {
	localAddr := common.GetAddr(tunnel.DstIp, tunnel.DstPort)
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalln(err)
	} else {
		localListener[localAddr] = &listener
	}
	go transport(
		client,
		common.GetAddr(tunnel.SrcIp, tunnel.SrcPort),
		localListener[localAddr],
	)
}

//  tunnel data to the remote server
func transport(client *ssh.Client, remoteAddr string, listener *net.Listener) {
	for {
		localConn, err := (*listener).Accept()
		if err != nil {
			log.Fatalln(err)
		}
		remoteConn, err := client.Dial("tcp", remoteAddr)
		if err != nil {
			log.Fatalln(err)
		}
		go io.Copy(localConn, remoteConn)
		go io.Copy(remoteConn, localConn)
	}
}

func main() {
	localListener = make(map[string]*net.Listener)
	sshClient, err := connected()
	if err != nil {
		log.Fatalln(err)
	}
	for _, tunnel := range tunnels {
		go setTunnel(sshClient, tunnel)

	}
	http.ListenAndServe(":6060", nil)
}
