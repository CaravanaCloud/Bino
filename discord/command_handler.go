package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

const UNKOWN_COMMAND_MESSAGE = "Desculpa, não entendi o seu comando"

var mentorshipCommand = commands.Mentorship{}

func processMessage(message *discordgo.MessageCreate) string {
	messageContent := message.Content
	if mentorshipCommand.CanHandle(messageContent) {
		response, _ := mentorshipCommand.Process(messageContent)
		return response
	}
	return UNKOWN_COMMAND_MESSAGE
}
