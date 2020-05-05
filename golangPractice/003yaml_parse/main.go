package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ContentDir", "../configs")
	viper.SetDefault("FileName", "config.yml")
	log.Println(ContentDir)
}
