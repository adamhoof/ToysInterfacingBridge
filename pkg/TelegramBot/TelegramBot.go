package TelegramBot

type TelegramBot interface {
	SetToken(pathToToken string)
	StartBot()
}
