package main

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/nats-io/stan.go"
	"l0-project/internal/model"
	"os"
)

func main() {
	natsConfig := os.Args[1:]
	sc, err := stan.Connect(natsConfig[0], natsConfig[1], stan.NatsURL(natsConfig[2]))
	if err != nil {
		fmt.Printf("Error during STAN connection: %s\n", err.Error())
		return
	}

	order := new(model.Order)

	_ = faker.SetRandomMapAndSliceSize(5)
	if err = faker.FakeData(&order); err != nil {
		fmt.Printf("Error during faking data with faker: %s\n", err.Error())
		return
	}

	fmt.Printf("%+v", order)

	bytes, err := json.Marshal(order)
	if err != nil {
		fmt.Printf("Error faking convert order to bytes[]: %s\n", err.Error())
		return
	}

	err = sc.Publish(natsConfig[3], bytes)
	if err != nil {
		fmt.Printf("Error publishing Order to nats: %s\n", err.Error())
	}

	if err := sc.Close(); err != nil {
		fmt.Printf("Error faking convert order to bytes[]: %s\n", err.Error())
		return
	}
}
