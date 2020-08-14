package config

import (
	"echo_sample/utils"
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Config ConfigのGetter
func Config() *viper.Viper {
	return config
}

// SetupConfig configを初期化する
func SetupConfig(mode string) {

	filepath := fmt.Sprintf("settings/config/%s.yml", mode)
	config = utils.YAMLToViper(filepath)

	config.Set("app.mode", mode)
}

// IsDevelopmentMode 開発モードか判定する
func IsDevelopmentMode() bool {
	return (config.GetString("app.mode") == "development")
}
