package app

import (
	"reflect"
	"testing"
)

func TestFeedbackHandler(t *testing.T) {
	handler := NewFeedbackHandler()

	if reflect.TypeOf(handler) != reflect.TypeOf(feedbackHandler{}) {
		t.Error("NewFeedbackHandler doesn't returns a feedbackHandler instance")
	}

	processCommandsTestcases := []struct {
		command       command
		expectedError error
	}{
		{
			command{text: "", userId: "UA01", userName: "slack.user"},
			errUsersNotFound,
		},
		{
			command{text: "Thanks <@UA01|slack.user>", userId: "UA01", userName: "slack.user"},
			errSelfFeedback,
		},
	}

	for _, tt := range processCommandsTestcases {
		_, err := handler.ProcessCommand(tt.command)

		if err != tt.expectedError {
			t.Error("ProcessCommand failed and processed a invalid command")
		}
	}

	validCommand := command{
		text:     "Thanks <@UA123|lucas.magnum> for helping me",
		userId:   "UA01",
		userName: "slack.user",
	}

	response, err := handler.ProcessCommand(validCommand)
	expectedResponse := handler.generateSuccessMessage([]user{
		user{"UA123", "lucas.magnum"},
	})

	if err != nil {
		t.Error("ProcessCommand failed to process a valid command")
	}

	if response != expectedResponse {
		t.Errorf(
			"ProcessCommand returned wrong results, expected: %v, got: %v",
			expectedResponse,
			response,
		)
	}
}
