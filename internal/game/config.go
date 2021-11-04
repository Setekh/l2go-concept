package game

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type ServerConfig struct {
	Hostname string
	Port     uint
}

type GeneralConfig struct {
	AuthPassword string
	ServerId     uint
}

type Configuration struct {
	Server  ServerConfig  `toml:"Server"`
	General GeneralConfig `toml:"General"`
}

var Config *Configuration

func init() {
	port, errParse := strconv.Atoi(os.Getenv("server.port"))

	if errParse != nil {
		port = 7777
	}

	var config = &Configuration{
		Server: ServerConfig{
			Hostname: os.Getenv("server.host"),
			Port:     uint(port),
		},
	}

	data, _ := ioutil.ReadFile("./config/server.toml")
	err := toml.Unmarshal(data, config)
	if err != nil {
		log.Println("Failed loading server config!", err.Error())
	}

	Config = config
}
