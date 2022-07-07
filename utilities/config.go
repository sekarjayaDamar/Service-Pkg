package utilities

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`

		MongoDb struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"mongo_db"`

		AuthService struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Endpoint struct {
				Verify string `yaml:"verify"`
			} `yaml:"endpoint"`
		} `yaml:"auth_service"`
	} `yaml:"server"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidateConfigPath(path string) *AppError {
	fileInfo, fileError := os.Stat(path)
	if fileError != nil {
		return NewLoadConfigError("Config Is a Directory, Not a Normal File")
	}
	if fileInfo.IsDir() {
		return NewLoadConfigError("Config Is a Directory, Not a Normal File")
	}
	return nil
}

func ParseFlags() *AppError {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if ValidateConfigPath(configPath) != nil {
		return NewLoadConfigError("Invalid Config Error, Cannot Load Config")
	}
	return nil
}
