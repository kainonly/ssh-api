package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"ssh-api/common"
	"ssh-api/testing"
	"sync"
	"time"
)

var (
	localListener map[string]*net.TCPListener
	tunnels       = []common.TunnelOption{
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
		{
			SrcIp:   "127.0.0.1",
			SrcPort: 6001,
			DstIp:   "127.0.0.1",
			DstPort: 6001,
		},
	}
)

func main() {
	localListener = make(map[string]*net.TCPListener)
	sshClient, err := testing.DebugConnected()
	if err != nil {
		log.Fatalln(err)
	}
	for _, tunnel := range tunnels {
		go setTunnel(sshClient, tunnel)
	}
	http.ListenAndServe(":6060", nil)
}

func setTunnel(client *ssh.Client, tunnel common.TunnelOption) {
	addr := common.GetAddr(tunnel.DstIp, tunnel.DstPort)
	remoteAddr := common.GetAddr(tunnel.SrcIp, tunnel.SrcPort)
	localAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	localListener[addr], err = net.ListenTCP("tcp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		localConn, err := localListener[addr].AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}
		go transport(localConn, client, remoteAddr)
	}
}

//  tunnel data to the remote server
func transport(localConn *net.TCPConn, client *ssh.Client, remoteAddr string) {
	defer localConn.Close()
	localConn.SetNoDelay(true)
	localConn.SetKeepAlive(true)
	localConn.SetKeepAlivePeriod(5 * time.Second)
	remoteConn, err := client.Dial("tcp", remoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer remoteConn.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		_, err := io.Copy(localConn, remoteConn)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		_, err := io.Copy(remoteConn, localConn)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
