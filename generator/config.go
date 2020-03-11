package generator

import "github.com/kelseyhightower/envconfig"

// Config contains generator configs
type Config struct {
	Count  int    `default:"1000"`
	Output string `default:"../example.csv"`
}

// GetExampleCSVGeneratorConfigFromEnv return example generator configs bases on environment variables
func GetExampleCSVGeneratorConfigFromEnv() (*Config, error) {
	c := new(Config)
	err := envconfig.Process("EXAMPLE", c)
	return c, err
}
