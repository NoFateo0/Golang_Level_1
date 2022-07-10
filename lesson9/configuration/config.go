package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

type Configuration struct {
	Port        int    `json:"port"`
	Database    string `json:"db_url"`
	Jaeger      string `json:"jaeger_url"`
	Sentry      string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppId   int    `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func FlagFunc() Configuration {
	parametersConf := Configuration{}

	flag.IntVar(&parametersConf.Port, "port", 8080, "Укажите нужный порт")
	flag.StringVar(&parametersConf.Database, "db_url", "", "Укажите ссылку на базу данных")
	flag.StringVar(&parametersConf.Jaeger, "jaeger_url", "", "Укажите ссылку на Jaeger")
	flag.StringVar(&parametersConf.Sentry, "sentry_url", "", "Укажите ссылку на Sentry")
	flag.StringVar(&parametersConf.KafkaBroker, "kafka_broker", "", "Укажите ссылку на Kafka")
	flag.IntVar(&parametersConf.SomeAppId, "some_app_id", 0, "Указать ID")
	flag.StringVar(&parametersConf.SomeAppKey, "some_app_key", "", "Указать ключ")

	flag.Parse()

	return parametersConf
}

func ConfFile() Configuration {
	file, err := os.Open("./package.json")
	if err != nil {
		log.Fatalf("Невозможно открыть файл: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Невозможно закрыть файл: %v", err)
		}
	}()

	packageJSON := Configuration{}

	err = json.NewDecoder(file).Decode(&packageJSON)
	if err != nil {
		log.Printf("Невозможно декодировать в структуру: %v", err)
	}

	return packageJSON
}

func Validation(configuration Configuration) Configuration {
	myIntConf := configuration
	var inInterface map[string]interface{}
	inMap := make(map[string]string)

	file, _ := json.Marshal(myIntConf)
	err := json.Unmarshal(file, &inInterface)
	if err != nil {
		log.Printf("Невозможно прочитать файл: %v", err)
	}

	for key, value := range inInterface {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		inMap[strKey] = strValue
	}

	for field, val := range inMap {
		switch field {
		case "jaeger_url":
			_, err := url.Parse(val)
			if err != nil {
				log.Printf("Невозможно получить URL: %v", err)
			}
		case "sentry_url":
			_, err := url.Parse(val)
			if err != nil {
				log.Printf("Невозможно получить URL: %v", err)
			}
		case "db_url":
			_, err := url.Parse(val)
			if err != nil {
				log.Printf("Невозможно получить URL: %v", err)
			}
		}
	}
	return myIntConf
}
