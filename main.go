package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/CaravanaCloud/bino/db"
	"github.com/CaravanaCloud/bino/discord"
	_ "github.com/CaravanaCloud/bino/locales"

	"github.com/bwmarrin/discordgo"
)

var Discord *discordgo.Session

func init() {
	var err error

	token := "Bot " + os.Getenv("BINO_DISCORD_TOKEN")

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
	Discord.AddHandler(discord.CommandMessageHandler)

	Discord.Identify.Intents = discordgo.IntentsGuildMessages

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Discord.Close()
}
