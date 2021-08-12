package commands

import (
	"testing"
)

func TestCanHandleListAllMentorsMessage(t *testing.T) {
	message := "mentorias"
	command := Mentorship{}

	if !command.CanHandle(message) {
		t.Fatalf("Expected mentorship to be able to handle message '%s'", message)
	}
}

func TestCanHandleListAllMessagesWithTraillingSpaces(t *testing.T) {
	message := "mentorias        "
	command := Mentorship{}

	if !command.CanHandle(message) {
		t.Fatalf("Expected mentorship to be able to handle message '%s'", message)
	}
}

func TestCanNotHandleOtherKindOfMessages(t *testing.T) {
	message := "ping"
	command := Mentorship{}

	if command.CanHandle(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}

func TestCanNotHandleMessagesWithMultipleWords(t *testing.T) {
	message := "ping mentorias"
	command := Mentorship{}

	if command.CanHandle(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}

func TestProcessListsAllMentors(t *testing.T) {
	message := "mentorias"
	command := Mentorship{}

	mentors, err := command.Process(message)
	expectedMentorList := "- Lucia\n- Julio\n- Marcus\n"

	if err != nil {
		t.Fatalf("Error processing '%s': %e", message, err)
	}

	if mentors != expectedMentorList {
		t.Fatalf("Expected mentors to be '%s', got: '%s'", expectedMentorList, mentors)
	}
}

func TestProcessShouldFailIfMessageCanNotBeHandled(t *testing.T) {
	message := "ping"
	command := Mentorship{}

	_, err := command.Process(message)

	if err == nil {
		t.Fatalf("Expected error while processing '%s'", message)
	}
}
