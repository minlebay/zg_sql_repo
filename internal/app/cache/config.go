package cache

import cfg "go.uber.org/config"

type Config struct {
	Address string `yaml:"address"`
	DB      string `yaml:"db"`
	ExpTime string `yaml:"exp_time"`
}

func NewCacheConfig(provider cfg.Provider) (*Config, error) {
	config := Config{}

	if err := provider.Get("cache").Populate(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
