package TelegramBot

type TelegramBot interface {
	SetToken(pathToToken string)
	StartBot()
	SendTextMessage(usr User, message string)
}
