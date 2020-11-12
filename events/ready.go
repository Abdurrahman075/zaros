package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func onReady(dc *discordgo.Session, r *discordgo.Ready) {
	log.Println("Login Sebagai " + r.User.Username + "#" + r.User.Discriminator)

	err := dc.UpdateStatus(0, "Ok")
	if err != nil {
		log.Println("Error saat ingin mengubah status")
	}
}
