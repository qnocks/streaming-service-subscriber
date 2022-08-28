package repository

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"l0-project/internal/model"
	"log"
	"testing"
)

func TestOrderRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	order := model.Order{OrderUid: "test"}
	bytes, _ := json.Marshal(order)

	expected := OrderEntity{
		OrderUID: "test",
		Data:     bytes,
	}

	mock.ExpectQuery("INSERT INTO orders").
		WithArgs(expected.OrderUID, expected.Data).
		WillReturnRows(&sqlmock.Rows{})

	repo := NewOrderRepository(db)
	repo.Save(expected)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s\n", err)
	}
}

func TestOrderRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	o1 := model.Order{OrderUid: "1"}
	o2 := model.Order{OrderUid: "2"}
	bytes1, _ := json.Marshal(o1)
	bytes2, _ := json.Marshal(o2)

	expected := []OrderEntity{
		{
			OrderUID: o1.OrderUid,
			Data:     bytes1,
		},
		{
			OrderUID: o2.OrderUid,
			Data:     bytes2,
		},
	}

	rows := sqlmock.NewRows([]string{"order_uid", "data"}).
		AddRow(expected[0].OrderUID, expected[0].Data).
		AddRow(expected[1].OrderUID, expected[1].Data)

	mock.ExpectQuery("SELECT (.+) FROM orders").WillReturnRows(rows)

	repo := NewOrderRepository(db)
	actual := repo.GetAll()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s\n", err)
	}

	assert.Equal(t, len(actual), 2)
	assert.Equal(t, actual[0].OrderUID, expected[0].OrderUID)
	assert.Equal(t, actual[1].OrderUID, expected[1].OrderUID)
}
