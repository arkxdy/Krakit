package config

type Config struct {
	Port     string
	Services Services
}

func Load() Config {
	return Config{
		Port:     ":8080",
		Services: LoadServices(),
	}
}
