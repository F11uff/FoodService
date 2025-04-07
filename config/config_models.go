package config

type Config struct {
	Env        string           `yaml:"env"`
	HttpServer HttpServerConfig `yaml:"http_server"`
	Db         DatabaseConfig   `yaml:"db"`
}

type HttpServerConfig struct {
	Address     string `yaml:"address"`
	Timeout     int    `yaml:"timeout"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}
