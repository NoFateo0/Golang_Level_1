package config

import (
	"flag"
	"fmt"
	"net/url"
	"strconv"
)

type Configuration struct {
	Port        int
	Database    *url.URL
	Jaeger      *url.URL
	Sentry      *url.URL
	KafkaBroker string
	SomeAppId   int
	SomeAppKey  string
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

func ParseInt(v string) (int, error) {
	valueInt, err := strconv.Atoi(v)

	if err != nil {
		return 0, fmt.Errorf("Невозможно прочитать передаваемое значение.")
	}

	return valueInt, nil
}

func ParseUrl(v string) (*url.URL, error) {
	valueURL, err := url.ParseRequestURI(v)

	if err != nil {
		return nil, fmt.Errorf("Невозможно прочитать передаваемое значение.")
	}

	return valueURL, nil
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
			data.Database, err = ParseUrl(v)
		case "jaeger_url":
			data.Jaeger, err = ParseUrl(v)
		case "sentry_url":
			data.Sentry, err = ParseUrl(v)
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
