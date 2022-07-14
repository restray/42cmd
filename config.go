package main

import (
	"log"
	"os"
	"path"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Oauth struct {
		Id     string `yaml:"id" envconfig:"OAUTH_CLIENT_ID"`
		Secret string `yaml:"secret" envconfig:"OAUTH_CLIENT_SECRET"`
	} `yaml:"oauth"`
}

func createConfigFolder() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configFolderPath := path.Join(dirname, ".42cmd")
	if err := os.Mkdir(configFolderPath, os.FileMode(int(0777))); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	return configFolderPath
}

func loadConfig() Config {
	var cfg Config

	loadConfigFile(&cfg)
	loadConfigEnv(&cfg)
	return cfg
}

func loadConfigFile(cfg *Config) {
	configPath := createConfigFolder()

	pathConfigFile := path.Join(configPath, "config.yml")
	f, err := os.Open(pathConfigFile)
	if err != nil {
		if os.IsNotExist(err) || os.IsPermission(err) {
			log.Printf("Config file %s do not exists, you should use it.\n", pathConfigFile)
			return
		} else {
			log.Fatal(err)
		}
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfigEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
	}
}
