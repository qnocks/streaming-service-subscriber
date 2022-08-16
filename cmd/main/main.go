package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error occurred during initialization config %s", err.Error())
	}

	fmt.Printf("Hello world! on port %d", viper.GetInt("port"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
