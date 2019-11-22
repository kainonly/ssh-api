package testing

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"ssh-api/client"
	"ssh-api/common"
)

type DebugOption struct {
	Host       string `json:"host"`
	Port       uint   `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Key        string `json:"key"`
	PassPhrase string `json:"passphrase"`
}

func GetDebugOption(filename string) (option common.ConnectOption, err error) {
	debug, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	var debugOption DebugOption
	if err = json.Unmarshal(debug, &debugOption); err != nil {
		return
	}
	key, err := base64.StdEncoding.DecodeString(debugOption.Key)
	if err != nil {
		return
	}
	option = common.ConnectOption{
		Host:       debugOption.Host,
		Port:       debugOption.Port,
		Username:   debugOption.Username,
		Password:   debugOption.Password,
		Key:        key,
		PassPhrase: []byte(debugOption.PassPhrase),
	}
	return
}

func DebugConnected() (sshClient *ssh.Client, err error) {
	option, err := GetDebugOption("./debug.json")
	if err != nil {
		return
	}
	c := client.InjectClient()
	return c.Testing(option)
}
