package stream

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"l0-project/internal/cache"
	"l0-project/internal/model"
	"l0-project/internal/repository"
	"log"
)

func Connect(clusterID, clientID, natsURL string) (stan.Conn, error) {
	return stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
}

type STAN struct {
	Conn  *stan.Conn
	Repo  *repository.OrderRepository
	Cache *cache.Cache
}

func NewSTAN(conn *stan.Conn, repo *repository.OrderRepository, cache *cache.Cache) *STAN {
	return &STAN{
		Conn:  conn,
		Repo:  repo,
		Cache: cache,
	}
}

func (s STAN) Listen(subject string) {
	_, err := (*s.Conn).Subscribe(subject, s.handleSubscribe, stan.StartWithLastReceived())

	if err != nil {
		log.Fatalf("Error listening to streaming connection: %s\n", err.Error())
	}
}

func (s STAN) handleSubscribe(msg *stan.Msg) {
	var data = msg.Data
	var order = *new(model.Order)
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		fmt.Printf("Error converting streamed bytes[] to order: %s\n", err.Error())
		return
	}

	validate := validator.New()
	err := validate.Struct(order)
	if err != nil {
		fmt.Printf("Error validating streamed order: %s\n", err.Error())
		return
	}

	s.Repo.Save(repository.OrderEntity{
		OrderUID: order.OrderUid,
		Data:     data,
	})

	s.Cache.Save(order)
}
