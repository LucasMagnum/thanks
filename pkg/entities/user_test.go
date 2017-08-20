package entities

import (
	"testing"
)

var user = User{
	Id:   "U205PLVJ5",
	Name: "lucas.magnum",
}

func TestUserEntity(t *testing.T) {
	if len(user.Id) == 0 {
		t.Error("User isn't exporting the Id")
	}

	if len(user.Name) == 0 {
		t.Error("User isn't exporting the Name")
	}
}
