package main

import (
	"log"
	"teashop/infrastructure"
	"teashop/router"
	"net/http"
)

// @title Swagger TeaHappy
// @verson 1.0
// @description API list for TeaHappy
// @host     localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.Println("Database name: ", infrastructure.GetDBName())
	log.Fatal(http.ListenAndServe(":"+infrastructure.GetAppPort(), router.Router()))
}
