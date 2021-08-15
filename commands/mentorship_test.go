package commands

import (
	"testing"
)

func TestCanRunListAllMentorsMessage(t *testing.T) {
	message := "mentorias"
	command := Mentorship{}

	if !command.CanRun(message) {
		t.Fatalf("Expected mentorship to be able to handle message '%s'", message)
	}
}

func TestCanRunListAllMessagesWithTraillingSpaces(t *testing.T) {
	message := "mentorias        "
	command := Mentorship{}

	if !command.CanRun(message) {
		t.Fatalf("Expected mentorship to be able to handle message '%s'", message)
	}
}

func TestCanNotRunOtherKindOfMessages(t *testing.T) {
	message := "ping"
	command := Mentorship{}

	if command.CanRun(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}

func TestCanNotRunMessagesWithMultipleWords(t *testing.T) {
	message := "ping mentorias"
	command := Mentorship{}

	if command.CanRun(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}

func TestRunListsAllMentors(t *testing.T) {
	message := "mentorias"
	command := Mentorship{}

	mentors, err := command.Run(message)
	expectedMentorList := "- Lucia\n- Julio\n- Marcus\n"

	if err != nil {
		t.Fatalf("Error processing '%s': %e", message, err)
	}

	if mentors != expectedMentorList {
		t.Fatalf("Expected mentors to be '%s', got: '%s'", expectedMentorList, mentors)
	}
}

func TestRunShouldFailIfMessageCanNotBeHandled(t *testing.T) {
	message := "ping"
	command := Mentorship{}

	_, err := command.Run(message)

	if err == nil {
		t.Fatalf("Expected error while processing '%s'", message)
	}
}
