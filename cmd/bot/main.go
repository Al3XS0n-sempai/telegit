package main

import (
	"fmt"

	"github.com/Al3XS0n-sempai/telegit/configs"
	"github.com/Al3XS0n-sempai/telegit/internal/bot"
)

func initConfigs() (*configs.CommonConfig, *configs.BotConfig) {
	fmt.Println("Reading CommonConfig & BotConfig")
	cfg, err := configs.GetCommonConfig("config.yaml")
	if err != nil {
		fmt.Printf("Error while reading CommonConfig %v", err)
		return nil, nil
	}
	botCfg, err := configs.GetBotConfig("config.yaml")
	if err != nil {
		fmt.Printf("Error while reading BotConfig %v", err)
		return nil, nil
	}

	fmt.Println("Validating CommonConfig")
	if err := cfg.Validate(); err != nil {
		fmt.Printf("Error while validating CommonConfig: %v", err)
		return nil, nil
	}
	if err := botCfg.Validate(); err != nil {
		fmt.Printf("Error while validating BotConfig: %v", err)
		return nil, nil
	}

	return cfg, botCfg
}

func settupLogger() {

}

func main() {
	cfg, botCfg := initConfigs()

	if cfg == nil || botCfg == nil {
		return
	}

	if cfg.MODE == configs.DEV_MODE {
		settupLogger()
	}

	bot, err := bot.GetNewTelegitBot(botCfg)

	if err != nil {
		fmt.Printf("Error while initializing TelegitBot %v", err)
		return
	}

	bot.SettupHandlers()

	fmt.Println("Starting telegram bot")
	if err := bot.Start(); err != nil {
		fmt.Println("Error while starting bot:\n", err)
	}
}
