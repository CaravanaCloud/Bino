package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

const UNKOWN_COMMAND_MESSAGE = "Desculpa, não entendi o seu comando"

func Handle(message *discordgo.MessageCreate) string {
	messageContent := message.Content
	if commands.CanHandle(messageContent) {
		response, _ := commands.Process(messageContent)
		return response
	}
	return UNKOWN_COMMAND_MESSAGE
}
