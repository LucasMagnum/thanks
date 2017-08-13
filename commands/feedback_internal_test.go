package commands

import (
    "testing"
    "reflect"
)


func TestParseUserData(t *testing.T) {
    tests := []struct{
        userData string
        expected []string
    }{
        {"<@123|user.name>", []string{"123", "user.name"}},
        {"<@123|user-name>", []string{"123", "user-name"}},
        {"@user.name", []string{"user.name", "user.name"}},
        {"@user-name", []string{"user-name", "user-name"}},
        {"@username", []string{"username", "username"}},
    }

    for _, tt := range tests {
        userId, username := parseUserData(tt.userData)

        if userId != tt.expected[0] || username != tt.expected[1] {
            t.Errorf(
                "parseUserData(%s): expected %v | got %v",
                tt.userData,
                tt.expected,
                []string{userId, username},
            )
        }
    }

}


func TestGetUsersFromText(t *testing.T) {
    tests := []struct{
        commandText string
        expected []string
    }{
        {"<@123|user-name> thanks <@124|username>", []string{"<@123|user-name>", "<@124|username>"}},
        {"Thanks <@123|user.name> <@124|user-name>", []string{"<@123|user.name>", "<@124|user-name>"}},
        {"@user.name thanks", []string{"@user.name"}},
        {"@username @username2 thanks", []string{"@username", "@username2"}},
        {"Thanks @user-name", []string{"@user-name"}},
        {"Thanks @user-name @username", []string{"@user-name", "@username"}},
    }

    for _, tt := range tests {
        users := getUsersFromText(tt.commandText)

        if !reflect.DeepEqual(users, tt.expected) {
            t.Errorf(
                "getUsersFromText(%s): expected %v | got %v",
                tt.commandText,
                tt.expected,
                users,
            )
        }
    }

}
