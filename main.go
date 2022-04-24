package main

import (
	"log"
	"net/http"

	"poe.market/config"
	"poe.market/streamreader"
	"poe.market/web"
)

func main() {
	configFile := config.NewConfigFile()

	go streamreader.LaunchReader(configFile)

	webHandler := web.NewHandler(configFile.PgConfig)

	if err := http.ListenAndServe(":8081", webHandler); err != nil {
		log.Fatal(err)
	}
}
