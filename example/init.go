package main

import (
	"golang.org/x/crypto/ssh"
	"net"
	"ssh-api/common"
)

var (
	localListener map[string]*net.Listener
	localConn     map[string]*net.Conn
	remoteConn    map[string]*net.Conn
)

var tunnels = []common.TunnelOption{
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

func connected() (client *ssh.Client, err error) {
	option, err := common.GetDebugOption("./debug.json")
	if err != nil {
		return
	}
	c := common.InjectClient()
	return c.Testing(option)
}
