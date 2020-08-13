package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

// YAMLToViper yamlファイルを読み込みviperインスタンスを初期化する
func YAMLToViper(filepath string) *viper.Viper {

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Errorf("ファイルオープンに失敗しました: filepaht=Z%s error=%s ", filepath, err.Error()))
	}

	// 環境変数解決
	expandContent := []byte(os.ExpandEnv(string(content)))

	v := viper.New()
	v.SetConfigType("yaml")
	readErr := v.ReadConfig(bytes.NewBuffer(expandContent))
	if readErr != nil {
		panic(fmt.Errorf("設定ファイルの読み込みに失敗しました: %s", err.Error()))
	}

	return v
}
