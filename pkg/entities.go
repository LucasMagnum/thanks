package app

type user struct {
	id   string
	name string
}

func (u *user) equal(otherUser user) bool {
	return u.id == otherUser.id && u.name == otherUser.name
}

type command struct {
	text     string
	userId   string
	userName string
}

func NewCommand(text string, userId string, userName string) command {
	return command{
		text:     text,
		userId:   userId,
		userName: userName,
	}
}

func NewUser(userId string, userName string) user {
	return user{
		id:   userId,
		name: userName,
	}
}
