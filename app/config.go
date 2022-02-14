package app

type Config struct {
	Level string
}

func NewConfig() (cfg *Config) {
	cfg = new(Config)
	cfg.Level = "info"
	return cfg
}
