package common

type Config struct {
	filePath string
	value    map[string]SSHOption
}

type SSHOption struct {
	Host       string `json:"host"`
	Port       uint   `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	PrivateKey string `json:"private_key"`
	Passphrase string `json:"passphrase"`
}

func (config *Config) init() {
	if config.value == nil {
		config.value = make(map[string]SSHOption)
	}
}

func (config *Config) Get(identity string) (exists bool, options SSHOption) {
	exists = config.value[identity] != (SSHOption{})
	options = config.value[identity]
	return
}

func (config *Config) Set(identity string, value *SSHOption) {
	config.value[identity] = *value
}
