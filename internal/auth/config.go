package auth

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type ServerConfig struct {
	Address string
	Port    int
}

type GeneralConfig struct {
	AuthPassword string
}

type Configuration struct {
	Server  ServerConfig  `toml:"Server"`
	General GeneralConfig `toml:"General"`
}

var Config *Configuration

func init() {
	var config = &Configuration{
		Server: ServerConfig{
			Address: "0.0.0.0",
			Port:    2106,
		},
		General: GeneralConfig{
			AuthPassword: "root",
		},
	}

	data, _ := ioutil.ReadFile("./config/Server.toml")
	err := toml.Unmarshal(data, config)

	if err != nil {
		fmt.Println("Failed loading config")
	}

	Config = config
}
