// Configs is a package that contains configurational types for Telegram bot and HttpServer
package configs

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// BotConfig is a struct that contains
type BotConfig struct {
	TOKEN           string        `yaml:"TELEGIT_BOT_TOKEN" env:"TELEGIT_BOT_TOKEN" env-default:"SOME_TOKEN"`
	TIMEOUT_SECONDS time.Duration `yaml:"TELEGIT_BOT_TIMEOUT_SECONDS" env:"TELEGIT_BOT_TIMEOUT_SECONDS" env-default:"10"`
}

func (cfg BotConfig) Validate() error {
	return nil
}

// GetBotConfig returns pointer to new instance of BotConfig initialized by config file
// configFile is the name of configurational file. It should lies in the same directory you are strating at.
func GetBotConfig(configFilename string) (*BotConfig, error) {
	cfg := &BotConfig{}

	if err := cleanenv.ReadConfig(configFilename, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Static check for interface satisfication
var _ Config = BotConfig{}
