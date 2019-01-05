package main

import (
	"log"

	"github.com/fsouza/go-dockerclient"
)

func listenDockerEvent() {
	client, err := docker.NewClient(Conf.DockerEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	listener := make(chan *docker.APIEvents)
	err = client.AddEventListener(listener)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = client.RemoveEventListener(listener)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for {
		select {
		case event := <-listener:
			if event.Action == "kill" {
				sendAlertToDiscord(event)
			}
		}
	}
}
