package TelegramBot

const meId = "558297691"

type User struct {
	id string
}

func (user *User) Recipient() string {
	return user.id
}

var me = User{id: meId}
