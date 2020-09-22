package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB struct {
		Driver string `json:"driver"`
		Addr   string `json:"addr"`
	} `json:"db"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Admin struct {
		Name     string `json:"name"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	} `json:"admin"`
	Address   string `json:"address"`
	Lang      string `json:"lang"`
	Secretkey string `json:"secretkey"`
}

var conf *Configuration

func Config() *Configuration {
	if conf != nil {
		return conf
	}

	var err error

	viper.SetConfigName("configuration")
	viper.AddConfigPath("/etc/guora/")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	if err = viper.Unmarshal(&conf); err != nil {
		fmt.Println("config file error:", err)
		os.Exit(1)
	}

	fmt.Println("Configuration.conf", conf)

	return conf
}
