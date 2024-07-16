package main

import (
	"log"

	"runners-postgresql/server"

	"runners-postgresql/config"

	_ "github.com/lib/pq"
)

/*
main.go should be as simple as possible
try to stick with initializing functions/methods and logs
*/

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing configuration")
	config := config.InitConfig("runners")
	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing Prometheus")
	go server.InitPrometheus()
	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
