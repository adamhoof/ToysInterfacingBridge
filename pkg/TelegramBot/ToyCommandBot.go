package TelegramBot

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"io/ioutil"
	"strings"
	"time"
)

type ToyCommandBot struct {
	bot *tb.Bot
}

func (toyCommandBot *ToyCommandBot) SetToken(pathToToken string) {

	token, err := ioutil.ReadFile(pathToToken)
	formattedToken := strings.Split(string(token), "\n")
	toyCommandBot.bot, err = tb.NewBot(tb.Settings{
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

func (toyCommandBot *ToyCommandBot) StartBot() {
	toyCommandBot.bot.Start()
}
