package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Discord *discordgo.Session

func init() {
	var err error

	token := "Bot " + os.Getenv("DISCORD_TOKEN")

	Discord, err = discordgo.New(token)
	if err != nil {
		log.Fatal("Wrong Token. Please check the token before connect to Discord")
	}

	err = Discord.Open()
	if err != nil {
		log.Fatal("Error creating Discord session", err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
}

func main() {

	Discord.AddHandler(appointmentMessage)

	// In this example, we only care about receiving message events.
	Discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Discord.Close()

}

func appointmentMessage(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if message.Author.ID == session.State.User.ID {
		return
	}

	fmt.Printf("%#v", message.Author)

	// If the message is "ping" reply with "Pong!"
	if message.Content == "agendar" {
		if session.State.User.Email == "marcus.pereira@live.com" {
			session.ChannelMessageSend(message.ChannelID, "Pong!")
		}
	}

	// If the message is "pong" reply with "Ping!"
	if message.Content == "pong" {
		session.ChannelMessageSend(message.ChannelID, "Ping!")
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
// func appointmentMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	toBino := strings.Contains(m.Content, "agendamento")

// 	if toBino && s.State.User.Email == "marcus.pereira@live.com" {
// 		s.ChannelMessageSend(m.ChannelID, "E cilada bino!!!")
// 		return
// 	} else {
// 		s.ChannelMessageSend(m.ChannelID, "Olá, somente mentores podem agendar mentorias. \nSeja um mentorado e marque a primeira conversa com a gente https://caravana.cloud/sessao-de-apresentacao")
// 		return
// 	}
// }
