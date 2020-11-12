package main

import (
	"github.com/bwmarrin/discordgo"
)

func AvatarCommands(dc *discordgo.Session, member *discordgo.Member, message *discordgo.Message, args []string) error {
	_, err := dc.ChannelMessageSend(message.ChannelID, member.User.AvatarURL("2048"))
	if err != nil {
		return err
	}

	return nil
}
