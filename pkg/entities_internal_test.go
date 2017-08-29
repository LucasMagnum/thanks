package app

import (
	"reflect"
	"testing"
)

func TestUser(t *testing.T) {
	userInstance := NewUser("01", "User 01")

	if reflect.TypeOf(userInstance) != reflect.TypeOf(user{}) {
		t.Error("NewUser doesn't returns an user instance")
	}

	if userInstance.id != "01" && userInstance.name != "User 01" {
		t.Error("New user doesn't created a properly user")
	}

	if !userInstance.equal(user{"01", "User 01"}) {
		t.Error("User equal should return true")
	}
}

func TestNewCommand(t *testing.T) {
	commandInstance := NewCommand("My command", "01", "User 01")

	if reflect.TypeOf(commandInstance) != reflect.TypeOf(command{}) {
		t.Error("NewCommand doesn't returns a command instance")
	}
}
