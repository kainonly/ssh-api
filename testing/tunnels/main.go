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
	"sync"
	"time"
)

var (
	kibana = client.TunnelOption{
		SrcIp:   "127.0.0.1",
		SrcPort: 5601,
		DstIp:   "127.0.0.1",
		DstPort: 5601,
	}
	service = client.TunnelOption{
		SrcIp:   "192.168.1.2",
		SrcPort: 8000,
		DstIp:   "127.0.0.1",
		DstPort: 8080,
	}
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	sshClient, err := testing.DebugConnected()
	if err != nil {
		log.Fatalln(err)
	}
	tunnel := &kibana
	remoteAddr := common.GetAddr(tunnel.SrcIp, tunnel.SrcPort)
	localAddr, err := net.ResolveTCPAddr("tcp", common.GetAddr(tunnel.DstIp, tunnel.DstPort))
	if err != nil {
		log.Fatalln(err)
	}
	localListener, err := net.ListenTCP("tcp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		localConn, err := localListener.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}
		go transport(
			localConn,
			sshClient,
			remoteAddr,
		)
	}
}

func transport(localConn *net.TCPConn, sshClient *ssh.Client, remoteAddr string) {
	defer localConn.Close()
	localConn.SetNoDelay(true)
	localConn.SetKeepAlive(true)
	localConn.SetKeepAlivePeriod(5 * time.Second)
	remoteConn, err := sshClient.Dial("tcp", remoteAddr)
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
