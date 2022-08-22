package cache

import (
	"encoding/json"
	"l0-project/internal/model"
	"l0-project/internal/repository"
	"l0-project/pkg/util"
)

type Cache struct {
	Orders []model.Order
}

func NewCache(repo repository.OrderRepository) *Cache {
	return &Cache{
		Orders: util.Map(repo.GetAll(), func(o repository.OrderEntity) model.Order {
			var order model.Order
			if err := json.Unmarshal(o.Data, &order); err != nil {
				return model.Order{}
			}

			return order
		}),
	}
}

func (c *Cache) Save(o model.Order) {
	c.Orders = append(c.Orders, o)
}

func (c Cache) GetAllOrders() []model.Order {
	return c.Orders
}

func (c Cache) GetOrderByID(id string) *model.Order {
	for _, o := range c.Orders {
		if o.OrderUid == id {
			return &o
		}
	}
	return nil
}
