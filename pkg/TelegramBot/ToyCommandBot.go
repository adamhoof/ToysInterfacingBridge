package TelegramBot

import (
	"fmt"
	tb "gopkg.in/telebot.v3"
	"io/ioutil"
	"strings"
	"time"
)

type ToyCommandBot struct {
	Bot *tb.Bot
}

func (toyCommandBot *ToyCommandBot) SetToken(pathToToken string) {

	token, err := ioutil.ReadFile(pathToToken)
	formattedToken := strings.Split(string(token), "\n")
	toyCommandBot.Bot, err = tb.NewBot(tb.Settings{
		Token: formattedToken[0],
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Bot token valid")
}

func (toyCommandBot *ToyCommandBot) StartBot() {
	toyCommandBot.Bot.Start()
}

func (toyCommandBot *ToyCommandBot) SendTextMessage(user User, message string) {
	_, err := toyCommandBot.Bot.Send(user, message)
	if err != nil {
		return
	}
}
