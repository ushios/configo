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

// String return string value.
func (c *Config) String(key string) (string, error) {
	return c.config.String(c.Environment, key)
}

// Int return int value.
func (c *Config) Int(key string) (int, error) {
	return c.config.Int(c.Environment, key)
}

// Floag return float value.
func (c *Config) Float(key string) (float64, error) {
	return c.config.Float(c.Environment, key)
}

// Bool return bool value.
func (c *Config) Bool(key string) (bool, error) {
	return c.config.Bool(c.Environment, key)
}
