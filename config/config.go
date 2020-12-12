package config

// Config represent app configuration
type Config struct {
	Port string
}

// GetConfig return configurations
func GetConfig() *Config {
	config := &Config{Port: "8080"}

	return config
}
