package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	filepath2 "path/filepath"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

func (cfg *DbConfig) GetDBInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User,
		cfg.Password, cfg.DBName, cfg.SSLMode)
}

func ParseDBConfig(filepath string) *DbConfig {
	fmt.Print(filepath2.Abs("./"))
	file, err := os.Open(filepath)
	if err != nil {
		slog.Error("Error during opening DBConfig: %v", err.Error())
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	parser := yaml.NewDecoder(file)
	var res DbConfig
	if err := parser.Decode(&res); err != nil {
		slog.Error("Error during parsing DBConfig: %v", err.Error())
		panic(err)
	}
	return &res
}
