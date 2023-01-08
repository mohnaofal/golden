package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mohnaofal/golden/microservice/topup-storage-service/config"
	"github.com/mohnaofal/golden/microservice/topup-storage-service/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// load config
	cfg := config.LoadConfig()

	// init router
	r := mux.NewRouter()

	go func() {
		// init event handler
		eventHandler := handler.NewEventHandler(cfg)
		eventHandler.RegisterConsumer()
	}()

	fmt.Println("Listening on port :", cfg.GetPORT())
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.GetPORT()), r); err != nil {
		log.Fatal(err)
	}
}
