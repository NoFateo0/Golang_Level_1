package main

import (
	"fmt"
	"lesson8/configuration"
)

func main() {
	confMapa := config.FlagFunc()
	conf, err := config.Parser(confMapa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)
}
