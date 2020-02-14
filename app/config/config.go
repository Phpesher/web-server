package Config

type Config struct {
	ListeningPort string
	Routes []string
}

// Init config
func InitConfig() *Config {
	conf := &Config{
		"8080",
		make([]string, 1),
	}

	return conf
}