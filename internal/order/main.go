package main

import (
	"fmt"

	"github.com/peileiscott/gorder/common/config"
	"github.com/spf13/viper"
)

func main() {
	if err := config.NewViperConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("order"))
}
