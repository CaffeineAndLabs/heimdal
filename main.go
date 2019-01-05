package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DockerEndpoint string `split_words:"true"`
	DiscordToken   string `split_words:"true"`
	DiscordChannel string `split_words:"true"`
}

var (
	Conf Config
)

func init() {
	err := envconfig.Process("HEIMDAL", &Conf)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Start")
	listenDockerEvent()
}
