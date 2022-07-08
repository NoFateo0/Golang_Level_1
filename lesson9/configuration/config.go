package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

func FlagFunc() map[string]string {
	var portFlag, databaseFlag, jaegerFlag, sentryFlag, kafkaBrokerFlag, someAppIdFlag, someAppKeyFlag string
	parametersMap := make(map[string]string)

	flag.StringVar(&portFlag, "port", "8080", "Укажите нужный порт")
	flag.StringVar(&databaseFlag, "db_url", "", "Укажите ссылку на базу данных")
	flag.StringVar(&jaegerFlag, "jaeger_url", "", "Укажите ссылку на Jaeger")
	flag.StringVar(&sentryFlag, "sentry_url", "", "Укажите ссылку на Sentry")
	flag.StringVar(&kafkaBrokerFlag, "kafka_broker", "", "Укажите ссылку на Kafka")
	flag.StringVar(&someAppIdFlag, "some_app_id", "0", "Указать ID")
	flag.StringVar(&someAppKeyFlag, "some_app_key", "", "Указать ключ")

	flag.Parse()

	flag.Visit(func(p *flag.Flag) {
		parametersMap[p.Name] = p.Value.String()
	})

	return parametersMap
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

func ParseInt(v string) (int, error) {
	valueInt, err := strconv.Atoi(v)

	if err != nil {
		return 0, fmt.Errorf("Невозможно прочитать передаваемое значение.")
	}

	return valueInt, nil
}

func ParseString(v string) (string, error) {
	if v == "" {
		err := fmt.Errorf("Невозможно прочитать передаваемое значение.")
		return v, err
	}
	return v, nil
}

func Parser(d map[string]string) (Configuration, error) {
	data := Configuration{}
	var err error

	for k, v := range d {
		switch k {
		case "port":
			data.Port, err = ParseInt(v)
		case "db_url":
			data.Database, err = ParseString(v)
		case "jaeger_url":
			data.Jaeger, err = ParseString(v)
		case "sentry_url":
			data.Sentry, err = ParseString(v)
		case "kafka_broker":
			data.KafkaBroker, err = ParseString(v)
		case "some_app_id":
			data.SomeAppId, err = ParseInt(v)
		case "some_app_key":
			data.SomeAppKey, err = ParseString(v)
		}
		if err != nil {
			return Configuration{}, err
		}
	}
	return data, nil
}
