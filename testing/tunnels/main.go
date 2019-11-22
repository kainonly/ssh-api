package main

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"ssh-api/common"
	"ssh-api/testing"
	"sync"
)

var (
	elastic = common.TunnelOption{
		SrcIp:   "127.0.0.1",
		SrcPort: 9200,
		DstIp:   "127.0.0.1",
		DstPort: 9200,
	}
	mysql = common.TunnelOption{
		SrcIp:   "127.0.0.1",
		SrcPort: 3306,
		DstIp:   "127.0.0.1",
		DstPort: 3306,
	}
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	common.InitBufPool()
	sshClient, err := testing.DebugConnected()
	if err != nil {
		log.Fatalln(err)
	}
	tunnel := &elastic
	remoteAddr := common.GetAddr(tunnel.SrcIp, tunnel.SrcPort)
	localAddr := common.GetAddr(tunnel.DstIp, tunnel.DstPort)
	localListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		remoteConn, err := sshClient.Dial("tcp", remoteAddr)
		if err != nil {
			log.Fatalln(err)
		}
		go forward(&localConn, &remoteConn)
	}
}

func forward(localConn *net.Conn, remoteConn *net.Conn) {
	defer (*localConn).Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		common.Copy(*localConn, *remoteConn)
	}()
	go func() {
		defer wg.Done()
		if _, err := common.Copy(*remoteConn, *localConn); err != nil {
			(*localConn).Close()
			(*remoteConn).Close()
		}
		(*remoteConn).Close()
	}()
	wg.Wait()
}
