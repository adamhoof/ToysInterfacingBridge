package TelegramBot

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"io/ioutil"
	"strings"
	"time"
)

type CommandBot struct {
	bot *tb.Bot
}

func (telegramBot *CommandBot) SetToken(pathToToken string) {

	token, err := ioutil.ReadFile(pathToToken)
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

func (telegramBot *CommandBot) StartBot() {
	telegramBot.bot.Start()
}
