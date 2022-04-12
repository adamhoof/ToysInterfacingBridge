package TelegramBot

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"io/ioutil"
	"strings"
	"time"
)

type TelegramBot struct {
	bot *tb.Bot
}

func (telegramBot *TelegramBot) CreateBot() {

	token, err := ioutil.ReadFile("Auth/BotToken")
	formattedToken := strings.Split(string(token), "\n")
	telegramBot.bot, err = tb.NewBot(tb.Settings{
		Token: formattedToken[0],
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("bot created")
}

func (telegramBot *TelegramBot) StartBot() {
	telegramBot.bot.Start()
}

func (telegramBot *TelegramBot) SendMessage(telegramBotHandler *TelegramBot, usr *User, title string, message interface{}) {

	_, err := telegramBotHandler.bot.Send(usr, title, message)
	if err != nil {
		fmt.Println("Failed to send message", err)
	}
}
