package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"MyDiscordBot/info"
)

var (
	token string
)

const (
	my_token_filename string = "MyDiscordBotToken"
)

func init() {
	token_file_path := os.Getenv("HOME") + "/" + my_token_filename

	file, err := os.Open(token_file_path)

	if err != nil {
		fmt.Println("Cannot open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	token = string(scanner.Text())
}

func main() {
	// Create discord session
	dg_session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("error creating Discord session:", err)
		return
	}

	// Register our message handler
	dg_session.AddHandler(message_handler)

	// Open websocket connection to Discord and begin listening.
	err = dg_session.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
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

	if m.Content[0] == '!' {
		cmd_arr := strings.Split(strings.Trim(m.Content, "!"), " ")

		switch cmd_arr[0] {
		case "ping":
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		case "pong":
			s.ChannelMessageSend(m.ChannelID, "pong!")
		case "help":
			s.ChannelMessageSend(m.ChannelID, info.Help_msg)
		default:
			s.ChannelMessageSend(m.ChannelID, "I do not understand.")
		}
	}
}
