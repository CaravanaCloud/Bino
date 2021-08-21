package commands

import (
	"testing"
)

func TestCanRunListAllMentorsMessage(t *testing.T) {
	message := "mentorias"

	if !mentorshipCommand.CanRun(message) {
		t.Fatalf("Expected mentorship to be able to handle message '%s'", message)
	}
}

func TestCanNotRunOtherKindOfMessages(t *testing.T) {
	message := "ping"

	if mentorshipCommand.CanRun(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}

func TestCanNotRunMessagesWithMultipleWords(t *testing.T) {
	message := "ping mentorias"

	if mentorshipCommand.CanRun(message) {
		t.Fatalf("Expected mentorship to not handle message '%s'", message)
	}
}

func TestRunListsAllMentors(t *testing.T) {
	message := "mentorias"

	mentors, err := mentorshipCommand.Run(message)
	expectedMentorList := "- Lucia\n- Julio\n- Marcus\n"

	if err != nil {
		t.Fatalf("Error processing '%s': %e", message, err)
	}

	if mentors != expectedMentorList {
		t.Fatalf("Expected mentors to be '%s', got: '%s'", expectedMentorList, mentors)
	}
}
