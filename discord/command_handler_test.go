package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestHandlesMentorshipCommandFromDiscordMessage(t *testing.T) {
	message := messageFromDiscordWith("mentorias")
	response := processMessage(message)
	expectedResponse := "- Lucia\n- Julio\n- Marcus\n"
	if response != expectedResponse {
		t.Fatalf("Expected message with content '%v' to be handled sucessfully", message.Content)
	}
}
func TestHandlesUnknownCommand(t *testing.T) {
	message := messageFromDiscordWith("unknown")
	response := processMessage(message)
	expectedResponse := "Desculpa, não entendi o seu comando"
	if response != expectedResponse {
		t.Fatalf("Expected message with content '%v' to be handled sucessfully", message.Content)
	}
}

func messageFromDiscordWith(content string) *discordgo.MessageCreate {
	message := discordgo.Message{Content: content}
	return &discordgo.MessageCreate{Message: &message}

}
