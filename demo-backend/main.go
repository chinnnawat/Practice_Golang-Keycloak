package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	// fmt.Println("Hello")
	initViper()
	title := viper.Get("Dashboard.title")
	fmt.Println(title)
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("demo")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		// fmt.Errorf = หากมี error ให้ show "unable to initialize viper: %w"
		// %w = ใช้ แสดง error หากเกิด error ขึ้น
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}
	log.Println("viper config initialized")
}
