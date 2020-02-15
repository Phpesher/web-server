package Config

type Config struct {
	ListeningPort string
}

// Init config
func InitConfig() *Config {
	conf := &Config{
		"8080",
	}

	return conf
}