package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const my_bot_client_token = "NzI5MzUwNTk0Nzk1NTM2NDA1.XwHqvA.S37s2UH0F2P9ahA1in7zc9b7ZEI"

func main() {
	// Create discord session
	dg_session, err := discordgo.New("Bot " + my_bot_client_token)

	if err != nil {
		fmt.Println("error creating Discord session: ", err)
		return
	}

	// Register our message handler
	dg_session.AddHandler(message_handler)

	// Open websocket connection to Discord and begin listening.
	err = dg_session.Open()
	if err != nil {
		fmt.Println("error opening connection: ", err)
	}

	fmt.Println("Bot is running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg_session.Close()
}

func message_handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Message is created by bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
