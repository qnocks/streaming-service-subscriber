package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"l0-project/pkg/model"
	"l0-project/pkg/repository"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error during initialization config: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error during loading env variables: %s\n", err.Error())
	}

	sc, _ := stan.Connect("test-cluster", "test", stan.NatsURL("0.0.0.0:4222"))

	var order model.Order
	var data []byte

	_, _ = sc.Subscribe("foo", func(msg *stan.Msg) {
		fmt.Printf("Recived %s\n", msg)
		data = msg.Data
		order = *new(model.Order)
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			fmt.Printf("Error during convert recieved bytes[] to order: %s\n", err.Error())
		}

		fmt.Printf("RECIEVED DATA:\n\n%v", order)

	}, stan.StartWithLastReceived())

	time.Sleep(1 * time.Second)

	//if err := sub.Unsubscribe(); err != nil {
	//	return
	//}

	//	===============================================================

	db, err := repository.NewDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		fmt.Printf("\n\nError during creating DB connection: %s\n", err.Error())
	}

	//const orderTable = "orders"

	queryJSON := fmt.Sprintf("INSERT INTO \"order\" (order_uid, data) VALUES ($1, $2)")

	queryRow := db.QueryRow(queryJSON, order.OrderUid, data)
	if queryRow.Err() != nil {
		fmt.Printf("Error during execution sql query: %s", queryRow.Err())
		return
	}

	//	===============================================================

	repo := repository.NewRepository()

	repo.LoadBackup(db)

	repo.Save(order)

	fmt.Println("LOCAL CACHE:")
	fmt.Println(repo.GetAll())

	//query := fmt.Sprintf("INSERT INTO %s ("+
	//	"order_uid,"+
	//	"track_number,"+
	//	"entry,"+
	//	"locale,"+
	//	"internal_signature,"+
	//	"customer_id,"+
	//	"delivery_service,"+
	//	"shardkey,"+
	//	"sm_id,"+
	//	"date_created,"+
	//	"oof_shard"+
	//	") values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
	//	orderTable)
	//row := db.QueryRow(query,
	//	order.OrderUid,
	//	order.TrackNumber,
	//	order.Entry,
	//	order.Locale,
	//	order.InternalSignature,
	//	order.CustomerID,
	//	order.DeliveryService,
	//	order.Shardkey,
	//	order.SmID,
	//	order.DateCreated,
	//	order.OofShard)
	//
	//if row.Err() != nil {
	//	fmt.Printf("Error during execution sql query: %s", row.Err())
	//}
	//
	//fmt.Println(row)

	router := gin.Default()
	router.GET("api/orders/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, repo.GetByID(id))
	})
	router.GET("api/orders", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, repo.GetAll())
	})

	if err = router.Run("localhost:8080"); err != nil {
		log.Fatalf("Error during running server: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
