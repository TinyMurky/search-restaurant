package config

// Config determine the struct thant <root>/configs/config.yaml should be
type Config struct {
	Server ServerConfig `yaml:"server"`
}

// ServerConfig is server related config
type ServerConfig struct {
	Port int `yaml:"port"`
}
