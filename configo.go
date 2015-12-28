package configo

import (
	"bufio"

	"github.com/ushios/config"
)

// Config has configurations.
type Config struct {
	Environment string
	config      *config.Config
}

// Create config.
func _config(reader *bufio.Reader, env string) (*Config, error) {
	cfg := new(Config)
	cfg.Environment = env

	ushiosConfig := config.NewDefault()
	err := ushiosConfig.UsingReader(reader)

	if err != nil {
		return cfg, err
	}

	cfg.config = ushiosConfig
	return cfg, err
}

// Instance return singleton instance.
func Instance(reader *bufio.Reader, env string) (*Config, error) {
	return _config(reader, env)
}

// Merge other cfg file.
func (c *Config) Merge(reader *bufio.Reader) error {
	source, err := _config(reader, c.Environment)
	if err != nil {
		return err
	}

	c.config.Merge(source.config)
	return nil
}

// String return string value.
func (c *Config) String(key string) (string, error) {
	return c.config.String(c.Environment, key)
}

// Int return int value.
func (c *Config) Int(key string) (int, error) {
	return c.config.Int(c.Environment, key)
}

// Float return float value.
func (c *Config) Float(key string) (float64, error) {
	return c.config.Float(c.Environment, key)
}

// Bool return bool value.
func (c *Config) Bool(key string) (bool, error) {
	return c.config.Bool(c.Environment, key)
}
