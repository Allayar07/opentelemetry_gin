package main

import (
	"log"
	"practice_optelem/second-service/internal/app/server"
)

func main() {
	if err := server.InitConfig(); err != nil {
		log.Fatalln(err)
	}
	server.Init("8081")
}
