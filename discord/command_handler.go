package discord

import "github.com/bwmarrin/discordgo"

func Handle(message *discordgo.MessageCreate) string {
	if message.Content == "mentorias" {
		return "- Lucia\n- Julio\n- Marcus\n"
	}
	return "Desculpa, não entendi o seu comando"
}
