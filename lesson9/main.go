package main

import (
	"fmt"
	"lesson9/configuration"
)

func main() {
	confMapa := config.FlagFunc()
	conf, err := config.Parser(confMapa)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonConf := config.ConfFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("С флагами: ", conf)
	fmt.Println("Из json: ", jsonConf)
}
