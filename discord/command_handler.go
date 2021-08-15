package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

const UNKOWN_COMMAND_MESSAGE = "Desculpa, não entendi o seu comando"

var mentorshipCommand = commands.Mentorship{}

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

func processMessage(messageContent string) string {
	if mentorshipCommand.CanRun(messageContent) {
		response, _ := mentorshipCommand.Run(messageContent)
		return response
	}
	return UNKOWN_COMMAND_MESSAGE
}
