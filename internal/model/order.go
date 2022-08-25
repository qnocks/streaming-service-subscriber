package model

type Order struct {
	OrderUid          string   `json:"order_uid" faker:"len=19" validate:"required"`
	TrackNumber       string   `json:"track_number" faker:"len=20" validate:"required"`
	Entry             string   `json:"entry" faker:"oneof: WBIL" validate:"required"`
	Delivery          Delivery `json:"delivery" validate:"required"`
	Payment           Payment  `json:"payment" validate:"required"`
	Items             []Item   `json:"items" validate:"required"`
	Locale            string   `json:"locale" faker:"oneof: en" validate:"required,len=2"`
	InternalSignature string   `json:"internal_signature" faker:"len=5" validate:"required"`
	CustomerID        string   `json:"customer_id" faker:"word" validate:"required"`
	DeliveryService   string   `json:"delivery_service" faker:"word" validate:"required"`
	Shardkey          string   `json:"shardkey" faker:"oneof: 9" validate:"required"`
	SmID              int64    `json:"sm_id" faker:"boundary_start=0, boundary_end=100" validate:"required"`
	DateCreated       string   `json:"date_created" faker:"date" validate:"required"`
	OofShard          string   `json:"oof_shard" faker:"oneof: 1" validate:"required"`
}

type Delivery struct {
	Name    string `json:"name" faker:"name" validate:"required"`
	Phone   string `json:"phone" faker:"len=11" validate:"required,len=11"`
	Zip     string `json:"zip" faker:"oneof: 2456454" validate:"required"`
	City    string `json:"city" faker:"word" validate:"required"`
	Address string `json:"address" faker:"word" validate:"required"`
	Region  string `json:"region" faker:"word" validate:"required"`
	Email   string `json:"email" faker:"email" validate:"required,email"`
}

type Payment struct {
	Transaction  string `json:"transaction" faker:"len=20" validate:"required"`
	RequestID    string `json:"request_id" faker:"len=20" validate:"required"`
	Currency     string `json:"currency" faker:"currency" validate:"required"`
	Provider     string `json:"provider" faker:"oneof: wbpay" validate:"required"`
	Amount       int64  `json:"amount" faker:"boundary_start=100, boundary_end=10000" validate:"required"`
	PaymentDt    int64  `json:"payment_dt" faker:"unix_time" validate:"required"`
	Bank         string `json:"bank" faker:"word" validate:"required"`
	DeliveryCost int64  `json:"delivery_cost" faker:"boundary_start=100, boundary_end=10000" validate:"required,gt=0"`
	GoodsTotal   int64  `json:"goods_total" faker:"boundary_start=1, boundary_end=100" validate:"required,gt=0"`
	CustomFee    int64  `json:"custom_fee" faker:"boundary_start=0, boundary_end=10000" validate:"required,gte=0"`
}

type Item struct {
	ChrtID      int64  `json:"chrt_id" faker:"boundary_start=100, boundary_end=10000" validate:"required"`
	TrackNumber string `json:"track_number" faker:"len=20" validate:"required"`
	Price       int64  `json:"price" faker:"boundary_start=100, boundary_end=10000" validate:"required,gte=0"`
	Rid         string `json:"rid" faker:"len=20" validate:"required"`
	Name        string `json:"name" faker:"first_name" validate:"required"`
	Sale        int64  `json:"sale" faker:"boundary_start=0, boundary_end=100" validate:"required,gte=0"`
	Size        string `json:"size" faker:"oneof: 1, 2, 5" validate:"required"`
	TotalPrice  int64  `json:"total_price" faker:"boundary_start=50, boundary_end=10000" validate:"required,gte=0"`
	NmID        int64  `json:"nm_id" faker:"boundary_start=1000, boundary_end=1000000" validate:"required"`
	Brand       string `json:"brand" faker:"word" validate:"required"`
	Status      int64  `json:"status" faker:"boundary_start=0, boundary_end=500" validate:"required"`
}
