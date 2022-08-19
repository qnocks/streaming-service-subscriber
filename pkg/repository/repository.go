package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"l0-project/pkg/model"
	"log"
)

type Repository struct {
	orders []model.Order
}

type OrderDTO struct {
	OrderUID string `db:"order_uid"`
	Data     []byte `db:"data"`
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Save(o model.Order) {
	r.orders = append(r.orders, o)
}

func (r Repository) GetAll() []model.Order {
	return r.orders
}

func (r *Repository) LoadBackup(db *sqlx.DB) {
	const orderTable = "order"
	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM \"%s\"", orderTable))
	if err != nil {
		log.Fatalf("Error during loading data to local cache: " + err.Error())
	}

	for rows.Next() {
		var o OrderDTO
		var order model.Order
		err := rows.StructScan(&o)
		if err != nil {
			log.Fatalf("\nError during mapping DB table to Order struct: " + err.Error())
		}

		json.Unmarshal(o.Data, &order)

		r.orders = append(r.orders, order)
	}

	fmt.Println("Hello world!")
}
