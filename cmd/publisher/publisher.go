package main

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/nats-io/stan.go"
	"l0-project/pkg/model"
)

func main() {
	sc, err := stan.Connect("test-cluster", "client-id", stan.NatsURL("0.0.0.0:4222"))
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
		fmt.Printf("Error during faking convert order to bytes[]: %s\n", err.Error())
		return
	}

	_ = sc.Publish("foo", bytes)

	if err := sc.Close(); err != nil {
		fmt.Printf("Error during faking convert order to bytes[]: %s\n", err.Error())
		return
	}
}
