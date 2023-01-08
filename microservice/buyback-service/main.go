package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mohnaofal/golden/microservice/buyback-service/config"
	"github.com/mohnaofal/golden/microservice/buyback-service/delivery"
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

	//  path prefix
	s := r.PathPrefix("/api").Subrouter()

	// init buyback delivery
	buybackDelivery := delivery.NewBuybackDelivery(cfg)
	buybackDelivery.Apply(s)

	fmt.Println("Listening on port :", cfg.GetPORT())
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.GetPORT()), r); err != nil {
		log.Fatal(err)
	}
}
