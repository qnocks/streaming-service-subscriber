# Streaming service subscriber

## Overview

The application subscribe to NATS streaming system to receiving orders and stores it in a database.

It exposes some API to get order by ID and simple frontend

[More information](docs/L0.pdf) 

## Usage

Create `app.env` `db.env` files with following configurations:

```
DB_USERNAME=
DB_NAME=
DB_HOST=
DB_PORT=
DB_PASSWORD=
DB_SSLMODE=
NATS_URL=
NATS_STREAMING_CLUSTER_ID=
NATS_STREAMING_CLIENT_ID=
NATS_STREAMING_SUBJECT=
```

```
POSTGRES_DB=
POSTGRES_USER=
POSTGRES_PASSWORD=
```

To run service run following:

```
$ docker-compose up
$ cd cmd/main
$ go build main.go
```

To run publisher script run following: 

```
$ cd cmd/publisher
$ go build <cluster_id> <client_id> <nats_url> <nats_subject>
```

## API

**GET** `/api/orders` - retrieving order by ID

---

**GET** `/` - basis frontend to use API

### Model

```json
{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```

