package common

type Config struct {
	FilePath string
	Value    map[string]string
}

func (config *Config) Get() map[string]string {
	return config.Value
}

func (config *Config) Set(value map[string]string) {
	config.Value = value
}
