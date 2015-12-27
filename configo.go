package configo

import "github.com/revel/config"

// singleton instance
var instance *Config

// Config has configurations.
type Config struct {
	Environment string
	config      *config.Config
}

// Create instance for singleton.
func newConfig(env string) (*Config, error) {
	cfg := new(Config)
	cfg.Environment = env
	revelConfig, err := config.ReadDefault("sample.cfg")

	if err != nil {
		return cfg, err
	}

	cfg.config = revelConfig
	return cfg, err
}

// Instance return singleton instance.
func Instance(env string) (*Config, error) {
	instance, err := newConfig(env)
	return instance, err
}

func (c *Config) String(key string) (string, error) {
	return c.config.String(c.Environment, key)
}
