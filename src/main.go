package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string = "MTA2MDE4MzkwMzk5NTk2NTQ1MQ.GAkDC3.hZds0Q-23smpaYGTwqqHyhkjKa4DPvfX3knQi8"
)


func init() {
	flag.StringVar(&Token, "t", Token, "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("⚠️ Could note found discord go")
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	// user, err := dg.User("@me")
	// if err != nil {
	// 	fmt.Println("⚠️ No user found")
	// 	return
	// }
	// fmt.Println(user.ID)

	<-make(chan struct{})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	// if m.Author.ID == s.State.User.ID {
	// 	return
	// }

	
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	// if m.Content == "pong" {
	// 	s.ChannelMessageSend(m.ChannelID, "Ping!")
	// }
}