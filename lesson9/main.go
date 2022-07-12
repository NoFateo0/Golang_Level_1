package main

import (
	"fmt"
	"lesson9/configuration"
)

func main() {
	confMapa := config.FlagFunc()
	gg := config.Validation(confMapa)

	jsonConf := config.ConfFile()
	gg2 := config.Validation(jsonConf)
	fmt.Println("С флагами: ", gg)
	fmt.Println("Из json: ", gg2)

}
