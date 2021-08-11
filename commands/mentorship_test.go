package commands

import (
	"testing"
)

func TestCanHandleListAllMentorsMessage(t *testing.T) {
	message := "mentorias"

	if !CanHandle(message) {
		t.Fatalf("Expected mentorship to be able to handle message '%s'", message)
	}
}

func TestCanHandleListAllMessagesWithTraillingSpaces(t *testing.T) {
	message := "mentorias        "

	if !CanHandle(message) {
		t.Fatalf("Expected mentorship to be able to handle message '%s'", message)
	}
}

func TestCanNotHandleOtherKindOfMessages(t *testing.T) {
	message := "ping"

	if CanHandle(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}

func TestCanNotHandleMessagesWithMultipleWords(t *testing.T) {
	message := "ping mentorias"

	if CanHandle(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}
