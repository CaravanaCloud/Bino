package discord

import (
	"testing"
)

func TestCommandMessageHandlerListsMentorships(t *testing.T) {
	message := "mentorias"
	response, _ := processMessage(message)
	expectedResponse := "- Lucia\n- Julio\n- Marcus\n"
	if response != expectedResponse {
		t.Fatalf("Expected message with content '%v' to be handled sucessfully", message)
	}
}
func TestCommandMessageHandlerNotifiesUserIfCommandIsUnknown(t *testing.T) {
	message := "unknown"
	response, _ := processMessage(message)
	expectedResponse := "Desculpa, não entendi o seu comando"
	if response != expectedResponse {
		t.Fatalf("Expected message with content '%v' to be handled sucessfully", message)
	}
}
