package configs

import (
	"errors"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

// Possible Startup modes
var modes map[string]int = map[string]int{
	"dev":  1,
	"prod": 2,
}

var (
	DEV_MODE  string = "dev"
	PROD_MODE string = "prod"
)

// CommonConfig contains variables that used for both Telegram Bot and Http Server
type CommonConfig struct {
	MODE string `yaml:"TELEGIT_MODE" env:"TELEGIT_MODE" env-default:"dev"`
}

func (cfg CommonConfig) Validate() error {
	if _, ok := modes[cfg.MODE]; !ok {
		return errors.New("TELEGIT_MODE should be one of the following:\n\tdev - for local development\n\tprod - for production")
	}

	return nil
}

// GetCommonConfig returns pointer to new instance of CommonConfig initialized by config file
// configFile is the name of configurational file. It should lies in the same directory you are strating at.
func GetCommonConfig(configFilename string) (*CommonConfig, error) {
	cfg := &CommonConfig{}

	if err := cleanenv.ReadConfig(configFilename, cfg); err != nil {
		return nil, err
	}

	cfg.MODE = strings.ToLower(cfg.MODE)

	return cfg, nil
}

// Static check for interface satisfication
var _ Config = CommonConfig{}
