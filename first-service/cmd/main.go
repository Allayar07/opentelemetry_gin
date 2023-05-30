package main

import (
	"log"
	"practice_optelem/first-service/internal/app/server"
	"practice_optelem/first-service/internal/configs"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatalln(err)
	}
	server.Init("8089")
}
