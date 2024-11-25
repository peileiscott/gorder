package main

import (
	"github.com/spf13/viper"
	"log"

	"github.com/peileiscott/gorder/common/config"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println(viper.Get("order"))
}
