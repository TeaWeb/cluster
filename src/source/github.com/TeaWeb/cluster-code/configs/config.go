package configs

import (
	"github.com/go-yaml/yaml"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/files"
	"github.com/iwind/TeaGo/logs"
)

var configFile = "config.yml"

type Config struct {
	Bind string `yaml:"bind"`
}

func SharedConfig() *Config {
	file := files.NewFile(Tea.ConfigFile(configFile))
	if !file.Exists() {
		return &Config{
			Bind: "0.0.0.0:7781",
		}
	}

	data, err := file.ReadAll()
	if err != nil {
		logs.Error(err)
		return &Config{
			Bind: "0.0.0.0:7781",
		}
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		logs.Error(err)
		return &Config{
			Bind: "0.0.0.0:7781",
		}
	}

	return config
}

func (this *Config) Save() error {
	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}

	file := files.NewFile(Tea.ConfigFile(configFile))
	return file.Write(data)
}
