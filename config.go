package quizzee_db

import (
	"bytes"
	"io/ioutil"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	if err = viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		panic(err)
	}
}

func ConfigString(key string) string {
	return viper.GetString(key)
}
