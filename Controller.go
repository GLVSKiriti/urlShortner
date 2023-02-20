package main

import (
	"log"

	"github.com/GLVSKiriti/urlShortner/config"
	"github.com/GLVSKiriti/urlShortner/handler"
	"github.com/GLVSKiriti/urlShortner/impl"
	_ "github.com/GLVSKiriti/urlShortner/storage"
	"github.com/valyala/fasthttp"
)

func main() {
	config, err := config.ReadFromFile("src/go.code/url-shortener/Config.json")

	if err != nil {
		log.Fatal("Can't find configuration. Error: %v", err)
	}

	redisClient, err := impl.NewPool(config.Redis.Host, config.Redis.Port, config.Redis.Password)

	if err != nil {
		log.Fatal("Could not connect to redis. Error: %v", err)
	}

	defer redisClient.Close()

	router := handler.New(config.Options.Schema, config.Options.Prefix, redisClient)

	log.Fatal(fasthttp.ListenAndServe(":"+config.Server.Port, router.Handler))

}
