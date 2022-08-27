package cache

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"l0-project/internal/model"
	"testing"
)

func newCacheInstance() *Cache {
	return &Cache{Orders: nil}
}

func TestCache_Save(t *testing.T) {
	const expectedOrdersSize = 1
	const expectedOrderUid = "test"
	cache := newCacheInstance()
	o := new(model.Order)
	o.OrderUid = "test"

	cache.Save(*o)

	assert.Equal(t, len(cache.Orders), expectedOrdersSize,
		fmt.Sprintf("FAILED Save(). Expected: %d, Got: %d", expectedOrdersSize, len(cache.Orders)))
	assert.Equal(t, cache.Orders[0].OrderUid, expectedOrderUid,
		fmt.Sprintf("FAILED Save(). Expected: %s, Got: %s", expectedOrderUid, cache.Orders[0].OrderUid))
}

func TestCache_GetOrderByID(t *testing.T) {
	expected := model.Order{OrderUid: "2"}
	orders := []model.Order{
		{OrderUid: "1"},
		expected,
		{OrderUid: "3"},
	}

	cache := newCacheInstance()
	cache.Save(orders[0])
	cache.Save(orders[1])
	cache.Save(orders[2])

	actual := cache.GetOrderByID("2")

	assert.Equal(t, *actual, expected,
		fmt.Sprintf("FAILED GetOrderByID(). Expected: %v, Got: %v", expected, *actual))
}
