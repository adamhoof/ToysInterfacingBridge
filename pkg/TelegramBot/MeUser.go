package TelegramBot

type MeUser struct {
	Id string
}

func (meUser *MeUser) Recipient() string {
	return meUser.Id
}
