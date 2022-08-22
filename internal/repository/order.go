package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	orderTable   = "orders"
	insertOrders = "INSERT INTO " + orderTable + " (order_uid, data) VALUES ($1, $2)"
	selectOrders = "SELECT * FROM " + orderTable
)

type OrderRepository struct {
	DB *sqlx.DB
}

type OrderEntity struct {
	OrderUID string `db:"order_uid"`
	Data     []byte `db:"data"`
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r OrderRepository) Save(orderEntity OrderEntity) {
	queryRow := r.DB.QueryRow(insertOrders, orderEntity.OrderUID, orderEntity.Data)
	if queryRow.Err() != nil {
		fmt.Printf("Error executing %s: %s", insertOrders, queryRow.Err())
		return
	}
}

func (r OrderRepository) GetAll() []OrderEntity {
	var orders []OrderEntity

	rows, err := r.DB.Queryx(selectOrders)
	if err != nil {
		return nil
	}

	for rows.Next() {
		var order OrderEntity
		err := rows.StructScan(&order)
		if err != nil {
			log.Fatalf("Error mapping DB table to struct: %s\n", err.Error())
		}

		orders = append(orders, order)
	}

	return orders
}
