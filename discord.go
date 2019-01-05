package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/fsouza/go-dockerclient"
)

func sendAlertToDiscord(event *docker.APIEvents) {
	dg, err := discordgo.New("Bot " + Conf.DiscordToken)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dg.Close()

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	message := formatMessage(event)
	_, err = dg.ChannelMessageSendComplex(Conf.DiscordChannel, message)
	if err != nil {
		log.Fatal(err)
	}
}

func formatMessage(event *docker.APIEvents) *discordgo.MessageSend {
	data := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Container action '%s' detected", event.Action),
	}

	msg := &discordgo.MessageSend{
		Content: fmt.Sprintf("%s has been killed", event.Actor.Attributes["name"]),
		Embed:   data,
	}

	return msg
}
