package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config данные конфигурации
type Config struct {
	Srv HTTPServer `yaml:"httpSrv"`
	DB  DB         `yaml:"db"`
}

// HTTPServer данные сервера HTTP
type HTTPServer struct {
	Addr string `yaml:"addr"`
}

// DB данные для подключения к БД
type DB struct {
	Conn string `yaml:"conn"`
}

// New получает данные конфигурации из файла
func New(path string) (*Config, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	return &cfg, yaml.NewDecoder(file).Decode(&cfg)
}
