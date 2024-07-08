package repository

import (
	cfg "go.uber.org/config"
)

type Db struct {
	DB       string `yaml:"database"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Dbs []Db `yaml:"dbs"`
}

func NewRepositoryConfig(provider cfg.Provider) (*Config, error) {
	config := Config{}
	dbs := []Db{}

	err := provider.Get("dbs").Populate(&dbs)
	if err != nil {
		return nil, err
	}
	config.Dbs = dbs

	return &config, nil
}
