package interactors

import (
	"testing"

	"github.com/lucasmagnum/thanks/pkg/entities"
)

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
