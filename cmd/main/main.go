package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"l0-project/internal"
	"l0-project/internal/cache"
	"l0-project/internal/handler"
	"l0-project/internal/repository"
	"l0-project/internal/stream"
	"l0-project/pkg/database/postgres"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s\n", err.Error())
	}

	if err := godotenv.Load("app.env"); err != nil {
		log.Fatalf("Error loading env variables: %s\n", err.Error())
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Error creating DB connection: %s\n", err.Error())
	}

	repo := repository.NewOrderRepository(db)
	c := cache.NewCache(*repo)

	conn, err := stream.Connect(
		os.Getenv("NATS_STREAMING_CLUSTER_ID"),
		os.Getenv("NATS_STREAMING_CLIENT_ID"),
		os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatalf("Error connecting to stan: %s\n", err.Error())
	}

	s := stream.NewSTAN(&conn, repo, c)
	s.Listen(os.Getenv("NATS_STREAMING_SUBJECT"))

	h := new(handler.Handler)
	h.InitRoutes(c)
	h.InitUI()

	server := new(internal.Server)
	err = server.Run(viper.GetString("server.port"), h.Router)
	if err != nil {
		log.Fatalf("Error running http server: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}
