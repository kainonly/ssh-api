package common

import (
	"strconv"
)

type ErrorMessage struct {
	error
}

// Get Addr
func GetAddr(ip string, port uint) string {
	return ip + ":" + strconv.Itoa(int(port))
}
