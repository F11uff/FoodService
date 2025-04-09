package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"pet/backend/internal/service"
)

func NewConfig(path string, logService service.LoggerService) (*Config, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		logService.Log.Info("Failed to open config file (config_constructor.go)")
		return nil, errors.New("File not exists")
	}

	decoder := yaml.NewDecoder(file)
	config := configConstructor()

	if err := decoder.Decode(&config); err != nil {
		logService.Log.Info("Failed to parse config file (config_constructor.go)")
		panic(err)
	}

	return config, nil
}

func configConstructor() *Config {
	return &Config{
		HttpServer: HttpServerConfig{
			"",
			0,
			0,
		},
		Db: DatabaseConfig{
			"",
			0,
			"",
			"",
			"",
		},
	}
}
