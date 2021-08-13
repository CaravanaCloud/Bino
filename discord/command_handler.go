package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

const UNKOWN_COMMAND_MESSAGE = "Desculpa, não entendi o seu comando"

var mentorshipCommand = commands.Mentorship{}

func CommandMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	response := processMessage(message.Content)
	session.ChannelMessageSend(message.ChannelID, response)

}

func processMessage(messageContent string) string {
	if mentorshipCommand.CanHandle(messageContent) {
		response, _ := mentorshipCommand.Process(messageContent)
		return response
	}
	return UNKOWN_COMMAND_MESSAGE
}
