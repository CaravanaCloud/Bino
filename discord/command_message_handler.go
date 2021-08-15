package discord

import (
	"github.com/CaravanaCloud/bino/commands"
	"github.com/bwmarrin/discordgo"
)

const unknownCommandMessage = "Desculpa, não entendi o seu comando"
const commandErrorMessage = "Desculpa, houve um erro no com seu comando"

var bino = Init(commands.Mentorship)

func CommandMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if messageIsFromBotItself(session, message) {
		return
	}

	response, err := processMessage(message.Content)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, commandErrorMessage)
	}

	session.ChannelMessageSend(message.ChannelID, response)

}

func messageIsFromBotItself(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	return message.Author.ID == session.State.User.ID
}

func processMessage(message string) (string, error) {
	return bino.Process(message), nil
}
