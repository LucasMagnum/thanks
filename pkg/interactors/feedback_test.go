package interactors

import (
	"reflect"
	"testing"

	"github.com/lucasmagnum/thanks/pkg/entities"
)

func TestFeedbackInteractor(t *testing.T) {
	interactor := FeedbackInteractor{}

	getUsersTestcases := []struct {
		text          string
		expectedUsers []entities.User
	}{
		{
			"Thank you <@UX01|lucasmagnum> for building this project",
			[]entities.User{
				entities.User{Id: "UX01", Name: "lucasmagnum"},
			},
		},
		{
			"<@UX02|lucas.magnum> <@UX04|lucas-magnum.01> Thanks",
			[]entities.User{
				entities.User{Id: "UX02", Name: "lucas.magnum"},
				entities.User{Id: "UX04", Name: "lucas-magnum.01"},
			},
		},
		{
			"Thanks <@UX03|lucas-magnum> <@UX03|lucas-magnum> <@UX03|lucas-magnum>",
			[]entities.User{
				entities.User{Id: "UX03", Name: "lucas-magnum"},
			},
		},
	}

	for _, tt := range getUsersTestcases {
		users := interactor.GetUsersFromText(tt.text)

		if !reflect.DeepEqual(users, tt.expectedUsers) {
			t.Errorf(
				"GetUsersFromText(%v): expected: %v, got: %v",
				tt.text,
				tt.expectedUsers,
				users,
			)
		}
	}

	validateUsersTestcases := []struct {
		user     entities.User
		users    []entities.User
		expected error
	}{
		{
			entities.User{Id: "UX01", Name: "lucasmagnum"},
			[]entities.User{
				entities.User{Id: "UX01", Name: "lucasmagnum"},
			},
			ErrSelfFeedback,
		},
		{
			entities.User{Id: "UX01", Name: "lucasmagnum"},
			[]entities.User{},
			ErrUsersNotFound,
		},
		{
			entities.User{Id: "UX03", Name: "zerothree"},
			[]entities.User{
				entities.User{Id: "UX02", Name: "lucas.magnum"},
				entities.User{Id: "UX04", Name: "lucas-magnum.01"},
			},
			nil,
		},
	}

	for _, tt := range validateUsersTestcases {
		err := interactor.ValidateUsers(tt.user, tt.users)

		if err != tt.expected {
			t.Errorf(
				"ValidateUsers(%v, %v): expected: %v, got: %v",
				tt.user,
				tt.users,
				tt.expected,
				err,
			)
		}
	}

}

// Test parseUserData
func TestparseUserData(t *testing.T) {
	testcases := []struct {
		userData         string
		expectedUserId   string
		expectedUserName string
	}{
		{"<@UX01|lucasmagnum>", "UX01", "lucasmagnum"},
		{"<@UX02|lucas.magnum>", "UX02", "lucas.magnum"},
		{"<@UX03|lucas-magnum>", "UX03", "lucas-magnum"},
		{"<@UX04|lucas-magnum.01>", "UX04", "lucas-magnum.01"},
	}

	for _, tt := range testcases {
		userId, userName := parseUserData(tt.userData)

		if userId != tt.expectedUserId {
			t.Errorf("Wrong userId, Expected: %v got: %v")
		}

		if userName != tt.expectedUserName {
			t.Errorf("Wrong userName, Expected: %v got: %v")
		}
	}
}

// Test findUser
func TestfindUser(t *testing.T) {
	users := []entities.User{
		entities.User{Id: "U101", Name: "user01"},
		entities.User{Id: "U102", Name: "user02"},
		entities.User{Id: "U103", Name: "user03"},
		entities.User{Id: "U104", Name: "user04"},
	}

	for _, user := range users {
		if !findUser(user, users) {
			t.Error("User should be found")
		}
	}
}
