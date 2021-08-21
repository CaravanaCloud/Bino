package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

var bino = Init(commands.Mentorship)

func CommandMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if messageIsFromBotItself(session, message) {
		return
	}

	response := bino.Process(message.Content)

	session.ChannelMessageSend(message.ChannelID, response)
}

func messageIsFromBotItself(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	return message.Author.ID == session.State.User.ID
}
