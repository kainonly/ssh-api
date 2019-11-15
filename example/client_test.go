package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"ssh-api/common"
	"testing"
)

type DebugOption struct {
	Host       string `json:"host"`
	Port       uint   `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Key        string `json:"key"`
	PassPhrase string `json:"passphrase"`
}

func GetOption() (option common.ConnectOption, err error) {
	debug, err := ioutil.ReadFile("../debug.json")
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

func TestSimple(t *testing.T) {
	option, err := GetOption()
	if err != nil {
		t.Error(err)
	}
	c := new(common.Client)
	client, err := c.Testing(option)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	println(string(client.ClientVersion()))
}

func TestPut(t *testing.T) {

}
