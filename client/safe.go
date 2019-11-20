package client

import (
	"net"
	"sync"
)

type safeMapListener struct {
	sync.RWMutex
	Map map[string]*net.TCPListener
}

func newSafeMapListener() *safeMapListener {
	listener := new(safeMapListener)
	listener.Map = make(map[string]*net.TCPListener)
	return listener
}

func (s *safeMapListener) Get(key string) *net.TCPListener {
	s.RLock()
	value := s.Map[key]
	s.RUnlock()
	return value
}

func (s *safeMapListener) Set(key string, listener *net.TCPListener) {
	s.Lock()
	s.Map[key] = listener
	s.Unlock()
}
