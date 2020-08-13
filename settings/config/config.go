package config

import (
	"echo_sample/utils"
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

// SetupConfig configを初期化する
func SetupConfig(mode string) {

	filepath := fmt.Sprintf("settings/config/%s.yml", mode)
	config = utils.YAMLToViper(filepath)
}

// Config ConfigのGetter
func Config() *viper.Viper {
	return config
}
