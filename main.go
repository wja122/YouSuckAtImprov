package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	//"net/http"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!youthere" {
		s.ChannelMessageSend(m.ChannelID, "Ben you suck")
		fmt.Println("Message I'm here sent")
	}
}

func main() {
	token := "MTA0ODM4NDExNjU0ODE5MDI1MA.GY_C69.oBBG_w_gWVyjwp9UBp9E_s-Lla8fM-unPrBHvg"
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	fmt.Println("Bot is running and ready. CTRL-c closes")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
