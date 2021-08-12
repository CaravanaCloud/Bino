package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

func Handle(message *discordgo.MessageCreate) string {
	messageContent := message.Content
	if commands.CanHandle(messageContent) {
		response, _ := commands.Process(messageContent)
		return response
	}
	return "Desculpa, não entendi o seu comando"
}
