package config

// DBConfig config for db
type DBConfig struct {
	Path string
}

// Config for system
type Config struct {
	DB   DBConfig
	Port string
}

// DefaultConfig get default
func DefaultConfig() Config {
	return Config{
		DB: DBConfig{
			Path: "./bin/db.db",
		},
		Port: ":2080",
	}
}
