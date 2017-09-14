package app

import (
	"reflect"
	"testing"
)

func TestFeedbackInteractor(t *testing.T) {
	interactor := feedbackInteractor{}

	validateCommandTestcases := []struct {
		command       command
		expectedError error
	}{
		{
			command{text: "", userId: "01", userName: "User01"},
			errUsersNotFound,
		},
		{
			command{text: "<@01|User01> thanks", userId: "01", userName: "User01"},
			errSelfFeedback,
		},
		{
			command{text: "Thanks <@02|User02>", userId: "01", userName: "User01"},
			nil,
		},
	}

	for _, tt := range validateCommandTestcases {
		if err := interactor.validateCommand(tt.command); err != tt.expectedError {
			t.Errorf(
				"validateCommand(%v): expected: %v, got: %v",
				tt.command,
				tt.expectedError,
				err,
			)
		}
	}

	parseUsersFromTextTestcases := []struct {
		text          string
		expectedUsers []user
	}{
		{
			"Thank you <@UX01|lucasmagnum> for building this project",
			[]user{
				user{id: "UX01", name: "lucasmagnum"},
			},
		},
		{
			"<@UX02|lucas.magnum> <@UX04|lucas-magnum.01> Thanks",
			[]user{
				user{id: "UX02", name: "lucas.magnum"},
				user{id: "UX04", name: "lucas-magnum.01"},
			},
		},
		{
			"Thanks <@UX03|lucas-magnum> <@UX03|lucas-magnum> <@UX03|lucas-magnum>",
			[]user{
				user{id: "UX03", name: "lucas-magnum"},
			},
		},
	}

	for _, tt := range parseUsersFromTextTestcases {
		users := interactor.parseUsersFromText(tt.text)

		if !reflect.DeepEqual(users, tt.expectedUsers) {
			t.Errorf(
				"parseUsersFromText(%v): expected: %v, got: %v",
				tt.text,
				tt.expectedUsers,
				users,
			)
		}
	}

	userDataTestcases := []struct {
		userData         string
		expectedUserId   string
		expectedUserName string
	}{
		{"<@UX01|lucasmagnum>", "UX01", "lucasmagnum"},
		{"<@UX02|lucas.magnum>", "UX02", "lucas.magnum"},
		{"<@UX03|lucas-magnum>", "UX03", "lucas-magnum"},
		{"<@UX04|lucas-magnum.01>", "UX04", "lucas-magnum.01"},
	}

	for _, tt := range userDataTestcases {
		userId, userName := interactor.parseUserData(tt.userData)

		if userId != tt.expectedUserId {
			t.Errorf("Wrong userId, Expected: %v got: %v")
		}

		if userName != tt.expectedUserName {
			t.Errorf("Wrong userName, Expected: %v got: %v")
		}
	}
}

func TestHasUser(t *testing.T) {
	user01 := NewUser("01", "User 01")

	users := []user{
		NewUser("02", "User 02"),
		NewUser("03", "User 02"),
		NewUser("04", "User 02"),
		NewUser("05", "User 02"),
	}

	if hasUser(users, user01) {
		t.Error("hasUser should return false, user01 isn't present on the users")
	}

	users = append(users, user01)

	if !hasUser(users, user01) {
		t.Error("hasUser should return true, user01 is present on the users")
	}
}
