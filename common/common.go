package common

import (
	"strconv"
)

type Error struct {
	Name string
	error
}

func SendError(identity string, err error) Error {
	return Error{
		Name:  identity,
		error: err,
	}
}

// Get Addr
func GetAddr(ip string, port uint) string {
	return ip + ":" + strconv.Itoa(int(port))
}
