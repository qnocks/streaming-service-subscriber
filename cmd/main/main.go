package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"l0-project/pkg/model"
	"log"
	"time"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error during initialization config: %s", err.Error())
		return
	}

	fmt.Printf("Hello world! on port: %d\n", viper.GetInt("port"))

	sc, _ := stan.Connect("test-cluster", "test", stan.NatsURL("0.0.0.0:4222"))

	sub, _ := sc.Subscribe("foo", func(msg *stan.Msg) {
		fmt.Printf("Recived %s\n", msg)
		order := new(model.Order)
		if err := json.Unmarshal(msg.Data, order); err != nil {
			fmt.Printf("Error during convert recieved bytes[] to order: %s\n", err.Error())
		}

		fmt.Printf("RECIEVED DATA:\n\n%v", order)

	}, stan.StartWithLastReceived())

	time.Sleep(1 * time.Second)

	if err := sub.Unsubscribe(); err != nil {
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
