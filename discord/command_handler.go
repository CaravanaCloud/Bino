package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

const UNKOWN_COMMAND_MESSAGE = "Desculpa, não entendi o seu comando"

func CommandMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if messageIsFromBotItself(session, message) {
		return
	}
	response := processMessage(message.Content)
	session.ChannelMessageSend(message.ChannelID, response)

}

func messageIsFromBotItself(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	return message.Author.ID == session.State.User.ID
}

func processMessage(message string) string {
	if message == "mentorias" {
		response, _ := commands.Mentorship.Run(message)
		return response
	}
	return UNKOWN_COMMAND_MESSAGE
}
