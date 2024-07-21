// internal/bot contains internal code of TelegitBot
package bot

import (
	"time"

	"github.com/Al3XS0n-sempai/telegit/configs"
	tele "gopkg.in/telebot.v3"
)

// TelegitBot is struct that contains all TelegitBot functionality
type TelegitBot struct {
	// cfg is TelegitBot config(API_KEY etc)
	cfg *configs.BotConfig

	// telebot is instance of telebot.v3.Bot
	telebot *tele.Bot
}

func (tb *TelegitBot) SettupHandlers() {
	tb.telebot.Handle("/test", TestHandler)
}

func (tb TelegitBot) Start() error {
	tb.telebot.Start()

	return nil
}

// GetNewTelegitBot returns new TelegitBot
func GetNewTelegitBot(cfg *configs.BotConfig) (*TelegitBot, error) {
	pref := tele.Settings{
		Token:  cfg.TOKEN,
		Poller: &tele.LongPoller{Timeout: time.Duration(cfg.TIMEOUT_SECONDS) * time.Second},
	}

	telebot, err := tele.NewBot(pref)

	if err != nil {
		return nil, err
	}

	return &TelegitBot{
		cfg:     cfg,
		telebot: telebot,
	}, nil
}
