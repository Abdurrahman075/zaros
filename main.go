package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := getEnv("DISCORD_TOKEN")

	dc, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	dc.AddHandler(onReady)
	dc.AddHandler(onMessageCreateCommandHandler)

	err = dc.Open()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dc.Close()
}
